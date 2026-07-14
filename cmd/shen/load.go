package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9643 := MakeNative(func(__e *ControlFlow) {
V893 := __e.Get(1)
_ = V893
tmp9644 := MakeNative(func(__e *ControlFlow) {
W894 := __e.Get(1)
_ = W894
tmp9645 := MakeNative(func(__e *ControlFlow) {
W895 := __e.Get(1)
_ = W895
tmp9646 := MakeNative(func(__e *ControlFlow) {
W901 := __e.Get(1)
_ = W901
__e.Return(symloaded)
return
}, 1)

var ifres9647 Obj

if True == W894 {
tmp9648 := Call(__e, PrimFunc(syminferences))


tmp9649 := Call(__e, PrimFunc(symshen_4app), tmp9648, MakeString(" inferences\n"), symshen_4a)


tmp9650 := PrimStringConcat(MakeString("\ntypechecked in "), tmp9649)

tmp9651 := Call(__e, PrimFunc(symstoutput))


tmp9652 := Call(__e, PrimFunc(sympr), tmp9650, tmp9651)


ifres9647 = tmp9652


} else {
ifres9647 = symshen_4skip


}

__e.TailApply(tmp9646, ifres9647)
return


}, 1)

tmp9653 := MakeNative(func(__e *ControlFlow) {
W896 := __e.Get(1)
_ = W896
tmp9654 := MakeNative(func(__e *ControlFlow) {
W897 := __e.Get(1)
_ = W897
tmp9655 := MakeNative(func(__e *ControlFlow) {
W898 := __e.Get(1)
_ = W898
tmp9656 := MakeNative(func(__e *ControlFlow) {
W899 := __e.Get(1)
_ = W899
tmp9657 := MakeNative(func(__e *ControlFlow) {
W900 := __e.Get(1)
_ = W900
__e.Return(W897)
return
}, 1)

tmp9658 := PrimStr(W899)

tmp9659 := PrimStringConcat(tmp9658, MakeString(" secs\n"))

tmp9660 := PrimStringConcat(MakeString("\nrun time: "), tmp9659)

tmp9661 := Call(__e, PrimFunc(symstoutput))


tmp9662 := Call(__e, PrimFunc(sympr), tmp9660, tmp9661)


__e.TailApply(tmp9657, tmp9662)
return


}, 1)

tmp9663 := PrimNumberSubtract(W898, W896)

__e.TailApply(tmp9656, tmp9663)
return


}, 1)

tmp9664 := PrimGetTime(symrun)

__e.TailApply(tmp9655, tmp9664)
return


}, 1)

tmp9665 := Call(__e, PrimFunc(symread_1file), V893)


tmp9666 := Call(__e, PrimFunc(symshen_4load_1help), W894, tmp9665)


__e.TailApply(tmp9654, tmp9666)
return


}, 1)

tmp9667 := PrimGetTime(symrun)

tmp9668 := Call(__e, tmp9653, tmp9667)


__e.TailApply(tmp9645, tmp9668)
return


}, 1)

tmp9669 := PrimValue(symshen_4_dtc_d)

__e.TailApply(tmp9644, tmp9669)
return


}, 1)

tmp9670 := Call(__e, ns2_1set, symload, tmp9643)


_ = tmp9670

tmp9671 := MakeNative(func(__e *ControlFlow) {
V904 := __e.Get(1)
_ = V904
V905 := __e.Get(2)
_ = V905
tmp9673 := PrimEqual(False, V904)

if True == tmp9673 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V905)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V905)
return
}


}, 2)

tmp9674 := Call(__e, ns2_1set, symshen_4load_1help, tmp9671)


_ = tmp9674

tmp9675 := MakeNative(func(__e *ControlFlow) {
V906 := __e.Get(1)
_ = V906
tmp9676 := MakeNative(func(__e *ControlFlow) {
Z907 := __e.Get(1)
_ = Z907
tmp9677 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z907)


tmp9678 := Call(__e, PrimFunc(symeval_1kl), tmp9677)


tmp9679 := Call(__e, PrimFunc(symshen_4app), tmp9678, MakeString("\n"), symshen_4s)


tmp9680 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9679, tmp9680)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp9676, V906)
return


}, 1)

tmp9681 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9675)


_ = tmp9681

tmp9682 := MakeNative(func(__e *ControlFlow) {
V908 := __e.Get(1)
_ = V908
tmp9683 := MakeNative(func(__e *ControlFlow) {
W909 := __e.Get(1)
_ = W909
tmp9684 := MakeNative(func(__e *ControlFlow) {
W911 := __e.Get(1)
_ = W911
tmp9685 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V908)
return
}, 0)

tmp9686 := MakeNative(func(__e *ControlFlow) {
Z913 := __e.Get(1)
_ = Z913
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z913, W909)
return
}, 1)

__e.TailApply(try_1catch, tmp9685, tmp9686)
return


}, 1)

tmp9687 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), W909)
return
}, 0)

tmp9688 := MakeNative(func(__e *ControlFlow) {
Z912 := __e.Get(1)
_ = Z912
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z912, W909)
return
}, 1)

tmp9689 := Call(__e, try_1catch, tmp9687, tmp9688)


__e.TailApply(tmp9684, tmp9689)
return


}, 1)

tmp9690 := MakeNative(func(__e *ControlFlow) {
Z910 := __e.Get(1)
_ = Z910
__e.TailApply(PrimFunc(symshen_4typetable), Z910)
return
}, 1)

tmp9691 := Call(__e, PrimFunc(symmapcan), tmp9690, V908)


__e.TailApply(tmp9683, tmp9691)
return


}, 1)

tmp9692 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9682)


_ = tmp9692

tmp9693 := MakeNative(func(__e *ControlFlow) {
V918 := __e.Get(1)
_ = V918
tmp9738 := PrimIsPair(V918)

var ifres9719 Obj

if True == tmp9738 {
tmp9736 := PrimHead(V918)

tmp9737 := PrimEqual(symdefine, tmp9736)

var ifres9721 Obj

if True == tmp9737 {
tmp9734 := PrimTail(V918)

tmp9735 := PrimIsPair(tmp9734)

var ifres9723 Obj

if True == tmp9735 {
tmp9731 := PrimTail(V918)

tmp9732 := PrimTail(tmp9731)

tmp9733 := PrimIsPair(tmp9732)

var ifres9725 Obj

if True == tmp9733 {
tmp9727 := PrimTail(V918)

tmp9728 := PrimTail(tmp9727)

tmp9729 := PrimHead(tmp9728)

tmp9730 := PrimEqual(sym_i, tmp9729)

var ifres9726 Obj

if True == tmp9730 {
ifres9726 = True


} else {
ifres9726 = False


}

ifres9725 = ifres9726


} else {
ifres9725 = False


}

var ifres9724 Obj

if True == ifres9725 {
ifres9724 = True


} else {
ifres9724 = False


}

ifres9723 = ifres9724


} else {
ifres9723 = False


}

var ifres9722 Obj

if True == ifres9723 {
ifres9722 = True


} else {
ifres9722 = False


}

ifres9721 = ifres9722


} else {
ifres9721 = False


}

var ifres9720 Obj

if True == ifres9721 {
ifres9720 = True


} else {
ifres9720 = False


}

ifres9719 = ifres9720


} else {
ifres9719 = False


}

if True == ifres9719 {
tmp9694 := PrimTail(V918)

tmp9695 := PrimHead(tmp9694)

tmp9696 := PrimTail(V918)

tmp9697 := PrimHead(tmp9696)

tmp9698 := PrimTail(V918)

tmp9699 := PrimTail(tmp9698)

tmp9700 := PrimTail(tmp9699)

tmp9701 := Call(__e, PrimFunc(symshen_4type_1F), tmp9697, tmp9700)


tmp9702 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9701)


tmp9703 := PrimCons(tmp9702, Nil)

__e.Return(PrimCons(tmp9695, tmp9703))
return


} else {
tmp9717 := PrimIsPair(V918)

var ifres9709 Obj

if True == tmp9717 {
tmp9715 := PrimHead(V918)

tmp9716 := PrimEqual(symdefine, tmp9715)

var ifres9711 Obj

if True == tmp9716 {
tmp9713 := PrimTail(V918)

tmp9714 := PrimIsPair(tmp9713)

var ifres9712 Obj

if True == tmp9714 {
ifres9712 = True


} else {
ifres9712 = False


}

ifres9711 = ifres9712


} else {
ifres9711 = False


}

var ifres9710 Obj

if True == ifres9711 {
ifres9710 = True


} else {
ifres9710 = False


}

ifres9709 = ifres9710


} else {
ifres9709 = False


}

if True == ifres9709 {
tmp9704 := PrimTail(V918)

tmp9705 := PrimHead(tmp9704)

tmp9706 := Call(__e, PrimFunc(symshen_4app), tmp9705, MakeString("\n"), symshen_4a)


tmp9707 := PrimStringConcat(MakeString("missing { in "), tmp9706)

__e.Return(PrimSimpleError(tmp9707))
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp9739 := Call(__e, ns2_1set, symshen_4typetable, tmp9693)


_ = tmp9739

tmp9740 := MakeNative(func(__e *ControlFlow) {
V925 := __e.Get(1)
_ = V925
V926 := __e.Get(2)
_ = V926
tmp9753 := PrimIsPair(V926)

var ifres9749 Obj

if True == tmp9753 {
tmp9751 := PrimHead(V926)

tmp9752 := PrimEqual(sym_j, tmp9751)

var ifres9750 Obj

if True == tmp9752 {
ifres9750 = True


} else {
ifres9750 = False


}

ifres9749 = ifres9750


} else {
ifres9749 = False


}

if True == ifres9749 {
__e.Return(Nil)
return
} else {
tmp9747 := PrimIsPair(V926)

if True == tmp9747 {
tmp9741 := PrimHead(V926)

tmp9742 := PrimTail(V926)

tmp9743 := Call(__e, PrimFunc(symshen_4type_1F), V925, tmp9742)


__e.Return(PrimCons(tmp9741, tmp9743))
return


} else {
tmp9744 := Call(__e, PrimFunc(symshen_4app), V925, MakeString("\n"), symshen_4a)


tmp9745 := PrimStringConcat(MakeString("missing } in "), tmp9744)

__e.Return(PrimSimpleError(tmp9745))
return


}


}


}, 2)

tmp9754 := Call(__e, ns2_1set, symshen_4type_1F, tmp9740)


_ = tmp9754

tmp9755 := MakeNative(func(__e *ControlFlow) {
V929 := __e.Get(1)
_ = V929
tmp9769 := PrimEqual(Nil, V929)

if True == tmp9769 {
__e.Return(Nil)
return
} else {
tmp9767 := PrimIsPair(V929)

var ifres9763 Obj

if True == tmp9767 {
tmp9765 := PrimTail(V929)

tmp9766 := PrimIsPair(tmp9765)

var ifres9764 Obj

if True == tmp9766 {
ifres9764 = True


} else {
ifres9764 = False


}

ifres9763 = ifres9764


} else {
ifres9763 = False


}

if True == ifres9763 {
tmp9756 := PrimHead(V929)

tmp9757 := PrimTail(V929)

tmp9758 := PrimHead(tmp9757)

tmp9759 := Call(__e, PrimFunc(symdeclare), tmp9756, tmp9758)


_ = tmp9759

tmp9760 := PrimTail(V929)

tmp9761 := PrimTail(tmp9760)

__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9761)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
return
}


}


}, 1)

tmp9770 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9755)


_ = tmp9770

tmp9771 := MakeNative(func(__e *ControlFlow) {
V934 := __e.Get(1)
_ = V934
V935 := __e.Get(2)
_ = V935
tmp9782 := PrimIsPair(V935)

var ifres9778 Obj

if True == tmp9782 {
tmp9780 := PrimTail(V935)

tmp9781 := PrimIsPair(tmp9780)

var ifres9779 Obj

if True == tmp9781 {
ifres9779 = True


} else {
ifres9779 = False


}

ifres9778 = ifres9779


} else {
ifres9778 = False


}

if True == ifres9778 {
tmp9772 := PrimHead(V935)

tmp9773 := Call(__e, PrimFunc(symdestroy), tmp9772)


_ = tmp9773

tmp9774 := PrimTail(V935)

tmp9775 := PrimTail(tmp9774)

__e.TailApply(PrimFunc(symshen_4unwind_1types), V934, tmp9775)
return


} else {
tmp9776 := PrimErrorToString(V934)

__e.Return(PrimSimpleError(tmp9776))
return


}


}, 2)

tmp9783 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9771)


_ = tmp9783

tmp9784 := MakeNative(func(__e *ControlFlow) {
V938 := __e.Get(1)
_ = V938
tmp9833 := PrimEqual(Nil, V938)

if True == tmp9833 {
__e.Return(Nil)
return
} else {
tmp9831 := PrimIsPair(V938)

var ifres9816 Obj

if True == tmp9831 {
tmp9829 := PrimTail(V938)

tmp9830 := PrimIsPair(tmp9829)

var ifres9818 Obj

if True == tmp9830 {
tmp9826 := PrimTail(V938)

tmp9827 := PrimTail(tmp9826)

tmp9828 := PrimIsPair(tmp9827)

var ifres9820 Obj

if True == tmp9828 {
tmp9822 := PrimTail(V938)

tmp9823 := PrimHead(tmp9822)

tmp9824 := PrimIntern(MakeString(":"))

tmp9825 := PrimEqual(tmp9823, tmp9824)

var ifres9821 Obj

if True == tmp9825 {
ifres9821 = True


} else {
ifres9821 = False


}

ifres9820 = ifres9821


} else {
ifres9820 = False


}

var ifres9819 Obj

if True == ifres9820 {
ifres9819 = True


} else {
ifres9819 = False


}

ifres9818 = ifres9819


} else {
ifres9818 = False


}

var ifres9817 Obj

if True == ifres9818 {
ifres9817 = True


} else {
ifres9817 = False


}

ifres9816 = ifres9817


} else {
ifres9816 = False


}

if True == ifres9816 {
tmp9785 := MakeNative(func(__e *ControlFlow) {
W939 := __e.Get(1)
_ = W939
tmp9801 := PrimEqual(W939, False)

if True == tmp9801 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9786 := MakeNative(func(__e *ControlFlow) {
W940 := __e.Get(1)
_ = W940
tmp9787 := MakeNative(func(__e *ControlFlow) {
W941 := __e.Get(1)
_ = W941
tmp9788 := PrimTail(V938)

tmp9789 := PrimTail(tmp9788)

tmp9790 := PrimTail(tmp9789)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9790)
return


}, 1)

tmp9791 := Call(__e, PrimFunc(symshen_4pretty_1type), W939)


tmp9792 := Call(__e, PrimFunc(symshen_4app), tmp9791, MakeString("\n"), symshen_4r)


tmp9793 := PrimStringConcat(MakeString(" : "), tmp9792)

tmp9794 := Call(__e, PrimFunc(symshen_4app), W940, tmp9793, symshen_4s)


tmp9795 := Call(__e, PrimFunc(symstoutput))


tmp9796 := Call(__e, PrimFunc(sympr), tmp9794, tmp9795)


__e.TailApply(tmp9787, tmp9796)
return


}, 1)

tmp9797 := PrimHead(V938)

tmp9798 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9797)


tmp9799 := Call(__e, PrimFunc(symeval_1kl), tmp9798)


__e.TailApply(tmp9786, tmp9799)
return


}


}, 1)

tmp9802 := PrimHead(V938)

tmp9803 := PrimTail(V938)

tmp9804 := PrimTail(tmp9803)

tmp9805 := PrimHead(tmp9804)

tmp9806 := Call(__e, PrimFunc(symshen_4typecheck), tmp9802, tmp9805)


__e.TailApply(tmp9785, tmp9806)
return


} else {
tmp9814 := PrimIsPair(V938)

if True == tmp9814 {
tmp9807 := PrimHead(V938)

tmp9808 := PrimIntern(MakeString(":"))

tmp9809 := PrimTail(V938)

tmp9810 := PrimCons(symA, tmp9809)

tmp9811 := PrimCons(tmp9808, tmp9810)

tmp9812 := PrimCons(tmp9807, tmp9811)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9812)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
return
}


}


}


}, 1)

tmp9834 := Call(__e, ns2_1set, symshen_4work_1through, tmp9784)


_ = tmp9834

tmp9835 := MakeNative(func(__e *ControlFlow) {
V943 := __e.Get(1)
_ = V943
tmp9977 := PrimIsPair(V943)

var ifres9851 Obj

if True == tmp9977 {
tmp9975 := PrimHead(V943)

tmp9976 := PrimIsPair(tmp9975)

var ifres9853 Obj

if True == tmp9976 {
tmp9972 := PrimHead(V943)

tmp9973 := PrimHead(tmp9972)

tmp9974 := PrimEqual(symlist, tmp9973)

var ifres9855 Obj

if True == tmp9974 {
tmp9969 := PrimHead(V943)

tmp9970 := PrimTail(tmp9969)

tmp9971 := PrimIsPair(tmp9970)

var ifres9857 Obj

if True == tmp9971 {
tmp9965 := PrimHead(V943)

tmp9966 := PrimTail(tmp9965)

tmp9967 := PrimTail(tmp9966)

tmp9968 := PrimEqual(Nil, tmp9967)

var ifres9859 Obj

if True == tmp9968 {
tmp9963 := PrimTail(V943)

tmp9964 := PrimIsPair(tmp9963)

var ifres9861 Obj

if True == tmp9964 {
tmp9960 := PrimTail(V943)

tmp9961 := PrimHead(tmp9960)

tmp9962 := PrimEqual(sym_1_1_6, tmp9961)

var ifres9863 Obj

if True == tmp9962 {
tmp9957 := PrimTail(V943)

tmp9958 := PrimTail(tmp9957)

tmp9959 := PrimIsPair(tmp9958)

var ifres9865 Obj

if True == tmp9959 {
tmp9953 := PrimTail(V943)

tmp9954 := PrimTail(tmp9953)

tmp9955 := PrimHead(tmp9954)

tmp9956 := PrimIsPair(tmp9955)

var ifres9867 Obj

if True == tmp9956 {
tmp9948 := PrimTail(V943)

tmp9949 := PrimTail(tmp9948)

tmp9950 := PrimHead(tmp9949)

tmp9951 := PrimHead(tmp9950)

tmp9952 := PrimEqual(symstr, tmp9951)

var ifres9869 Obj

if True == tmp9952 {
tmp9943 := PrimTail(V943)

tmp9944 := PrimTail(tmp9943)

tmp9945 := PrimHead(tmp9944)

tmp9946 := PrimTail(tmp9945)

tmp9947 := PrimIsPair(tmp9946)

var ifres9871 Obj

if True == tmp9947 {
tmp9937 := PrimTail(V943)

tmp9938 := PrimTail(tmp9937)

tmp9939 := PrimHead(tmp9938)

tmp9940 := PrimTail(tmp9939)

tmp9941 := PrimHead(tmp9940)

tmp9942 := PrimIsPair(tmp9941)

var ifres9873 Obj

if True == tmp9942 {
tmp9930 := PrimTail(V943)

tmp9931 := PrimTail(tmp9930)

tmp9932 := PrimHead(tmp9931)

tmp9933 := PrimTail(tmp9932)

tmp9934 := PrimHead(tmp9933)

tmp9935 := PrimHead(tmp9934)

tmp9936 := PrimEqual(symlist, tmp9935)

var ifres9875 Obj

if True == tmp9936 {
tmp9923 := PrimTail(V943)

tmp9924 := PrimTail(tmp9923)

tmp9925 := PrimHead(tmp9924)

tmp9926 := PrimTail(tmp9925)

tmp9927 := PrimHead(tmp9926)

tmp9928 := PrimTail(tmp9927)

tmp9929 := PrimIsPair(tmp9928)

var ifres9877 Obj

if True == tmp9929 {
tmp9915 := PrimTail(V943)

tmp9916 := PrimTail(tmp9915)

tmp9917 := PrimHead(tmp9916)

tmp9918 := PrimTail(tmp9917)

tmp9919 := PrimHead(tmp9918)

tmp9920 := PrimTail(tmp9919)

tmp9921 := PrimTail(tmp9920)

tmp9922 := PrimEqual(Nil, tmp9921)

var ifres9879 Obj

if True == tmp9922 {
tmp9909 := PrimTail(V943)

tmp9910 := PrimTail(tmp9909)

tmp9911 := PrimHead(tmp9910)

tmp9912 := PrimTail(tmp9911)

tmp9913 := PrimTail(tmp9912)

tmp9914 := PrimIsPair(tmp9913)

var ifres9881 Obj

if True == tmp9914 {
tmp9902 := PrimTail(V943)

tmp9903 := PrimTail(tmp9902)

tmp9904 := PrimHead(tmp9903)

tmp9905 := PrimTail(tmp9904)

tmp9906 := PrimTail(tmp9905)

tmp9907 := PrimTail(tmp9906)

tmp9908 := PrimEqual(Nil, tmp9907)

var ifres9883 Obj

if True == tmp9908 {
tmp9898 := PrimTail(V943)

tmp9899 := PrimTail(tmp9898)

tmp9900 := PrimTail(tmp9899)

tmp9901 := PrimEqual(Nil, tmp9900)

var ifres9885 Obj

if True == tmp9901 {
tmp9887 := PrimHead(V943)

tmp9888 := PrimTail(tmp9887)

tmp9889 := PrimHead(tmp9888)

tmp9890 := PrimTail(V943)

tmp9891 := PrimTail(tmp9890)

tmp9892 := PrimHead(tmp9891)

tmp9893 := PrimTail(tmp9892)

tmp9894 := PrimHead(tmp9893)

tmp9895 := PrimTail(tmp9894)

tmp9896 := PrimHead(tmp9895)

tmp9897 := PrimEqual(tmp9889, tmp9896)

var ifres9886 Obj

if True == tmp9897 {
ifres9886 = True


} else {
ifres9886 = False


}

ifres9885 = ifres9886


} else {
ifres9885 = False


}

var ifres9884 Obj

if True == ifres9885 {
ifres9884 = True


} else {
ifres9884 = False


}

ifres9883 = ifres9884


} else {
ifres9883 = False


}

var ifres9882 Obj

if True == ifres9883 {
ifres9882 = True


} else {
ifres9882 = False


}

ifres9881 = ifres9882


} else {
ifres9881 = False


}

var ifres9880 Obj

if True == ifres9881 {
ifres9880 = True


} else {
ifres9880 = False


}

ifres9879 = ifres9880


} else {
ifres9879 = False


}

var ifres9878 Obj

if True == ifres9879 {
ifres9878 = True


} else {
ifres9878 = False


}

ifres9877 = ifres9878


} else {
ifres9877 = False


}

var ifres9876 Obj

if True == ifres9877 {
ifres9876 = True


} else {
ifres9876 = False


}

ifres9875 = ifres9876


} else {
ifres9875 = False


}

var ifres9874 Obj

if True == ifres9875 {
ifres9874 = True


} else {
ifres9874 = False


}

ifres9873 = ifres9874


} else {
ifres9873 = False


}

var ifres9872 Obj

if True == ifres9873 {
ifres9872 = True


} else {
ifres9872 = False


}

ifres9871 = ifres9872


} else {
ifres9871 = False


}

var ifres9870 Obj

if True == ifres9871 {
ifres9870 = True


} else {
ifres9870 = False


}

ifres9869 = ifres9870


} else {
ifres9869 = False


}

var ifres9868 Obj

if True == ifres9869 {
ifres9868 = True


} else {
ifres9868 = False


}

ifres9867 = ifres9868


} else {
ifres9867 = False


}

var ifres9866 Obj

if True == ifres9867 {
ifres9866 = True


} else {
ifres9866 = False


}

ifres9865 = ifres9866


} else {
ifres9865 = False


}

var ifres9864 Obj

if True == ifres9865 {
ifres9864 = True


} else {
ifres9864 = False


}

ifres9863 = ifres9864


} else {
ifres9863 = False


}

var ifres9862 Obj

if True == ifres9863 {
ifres9862 = True


} else {
ifres9862 = False


}

ifres9861 = ifres9862


} else {
ifres9861 = False


}

var ifres9860 Obj

if True == ifres9861 {
ifres9860 = True


} else {
ifres9860 = False


}

ifres9859 = ifres9860


} else {
ifres9859 = False


}

var ifres9858 Obj

if True == ifres9859 {
ifres9858 = True


} else {
ifres9858 = False


}

ifres9857 = ifres9858


} else {
ifres9857 = False


}

var ifres9856 Obj

if True == ifres9857 {
ifres9856 = True


} else {
ifres9856 = False


}

ifres9855 = ifres9856


} else {
ifres9855 = False


}

var ifres9854 Obj

if True == ifres9855 {
ifres9854 = True


} else {
ifres9854 = False


}

ifres9853 = ifres9854


} else {
ifres9853 = False


}

var ifres9852 Obj

if True == ifres9853 {
ifres9852 = True


} else {
ifres9852 = False


}

ifres9851 = ifres9852


} else {
ifres9851 = False


}

if True == ifres9851 {
tmp9836 := PrimTail(V943)

tmp9837 := PrimTail(tmp9836)

tmp9838 := PrimHead(tmp9837)

tmp9839 := PrimTail(tmp9838)

tmp9840 := PrimHead(tmp9839)

tmp9841 := PrimTail(V943)

tmp9842 := PrimTail(tmp9841)

tmp9843 := PrimHead(tmp9842)

tmp9844 := PrimTail(tmp9843)

tmp9845 := PrimTail(tmp9844)

tmp9846 := PrimCons(sym_a_a_6, tmp9845)

__e.Return(PrimCons(tmp9840, tmp9846))
return


} else {
tmp9849 := PrimIsPair(V943)

if True == tmp9849 {
tmp9847 := MakeNative(func(__e *ControlFlow) {
Z944 := __e.Get(1)
_ = Z944
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z944)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9847, V943)
return


} else {
__e.Return(V943)
return
}


}


}, 1)

tmp9978 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9835)


_ = tmp9978

tmp9979 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("type error\n")))
return
}, 0)

tmp9980 := Call(__e, ns2_1set, symshen_4type_1error, tmp9979)


_ = tmp9980

tmp9981 := MakeNative(func(__e *ControlFlow) {
V945 := __e.Get(1)
_ = V945
tmp9982 := MakeNative(func(__e *ControlFlow) {
W946 := __e.Get(1)
_ = W946
tmp9983 := MakeNative(func(__e *ControlFlow) {
W947 := __e.Get(1)
_ = W947
tmp9984 := MakeNative(func(__e *ControlFlow) {
W948 := __e.Get(1)
_ = W948
tmp9985 := MakeNative(func(__e *ControlFlow) {
W949 := __e.Get(1)
_ = W949
tmp9986 := MakeNative(func(__e *ControlFlow) {
W951 := __e.Get(1)
_ = W951
__e.Return(W946)
return
}, 1)

tmp9987 := Call(__e, PrimFunc(symshen_4write_1kl), W949, W948)


__e.TailApply(tmp9986, tmp9987)
return


}, 1)

tmp9988 := MakeNative(func(__e *ControlFlow) {
Z950 := __e.Get(1)
_ = Z950
tmp9989 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z950)


__e.TailApply(PrimFunc(symshen_4partial), tmp9989)
return


}, 1)

tmp9990 := Call(__e, PrimFunc(symmap), tmp9988, W947)


__e.TailApply(tmp9985, tmp9990)
return


}, 1)

tmp9991 := PrimOpenStream(W946, symout)

__e.TailApply(tmp9984, tmp9991)
return


}, 1)

tmp9992 := Call(__e, PrimFunc(symread_1file), V945)


__e.TailApply(tmp9983, tmp9992)
return


}, 1)

tmp9993 := Call(__e, PrimFunc(symshen_4klfile), V945)


__e.TailApply(tmp9982, tmp9993)
return


}, 1)

tmp9994 := Call(__e, ns2_1set, symbootstrap, tmp9981)


_ = tmp9994

tmp9995 := MakeNative(func(__e *ControlFlow) {
V952 := __e.Get(1)
_ = V952
tmp10018 := PrimIsPair(V952)

var ifres10005 Obj

if True == tmp10018 {
tmp10016 := PrimHead(V952)

tmp10017 := PrimEqual(symshen_4f_1error, tmp10016)

var ifres10007 Obj

if True == tmp10017 {
tmp10014 := PrimTail(V952)

tmp10015 := PrimIsPair(tmp10014)

var ifres10009 Obj

if True == tmp10015 {
tmp10011 := PrimTail(V952)

tmp10012 := PrimTail(tmp10011)

tmp10013 := PrimEqual(Nil, tmp10012)

var ifres10010 Obj

if True == tmp10013 {
ifres10010 = True


} else {
ifres10010 = False


}

ifres10009 = ifres10010


} else {
ifres10009 = False


}

var ifres10008 Obj

if True == ifres10009 {
ifres10008 = True


} else {
ifres10008 = False


}

ifres10007 = ifres10008


} else {
ifres10007 = False


}

var ifres10006 Obj

if True == ifres10007 {
ifres10006 = True


} else {
ifres10006 = False


}

ifres10005 = ifres10006


} else {
ifres10005 = False


}

if True == ifres10005 {
tmp9996 := PrimTail(V952)

tmp9997 := PrimHead(tmp9996)

tmp9998 := PrimStr(tmp9997)

tmp9999 := PrimStringConcat(MakeString("partial function "), tmp9998)

tmp10000 := PrimCons(tmp9999, Nil)

__e.Return(PrimCons(symsimple_1error, tmp10000))
return


} else {
tmp10003 := PrimIsPair(V952)

if True == tmp10003 {
tmp10001 := MakeNative(func(__e *ControlFlow) {
Z953 := __e.Get(1)
_ = Z953
__e.TailApply(PrimFunc(symshen_4partial), Z953)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp10001, V952)
return


} else {
__e.Return(V952)
return
}


}


}, 1)

tmp10019 := Call(__e, ns2_1set, symshen_4partial, tmp9995)


_ = tmp10019

tmp10020 := MakeNative(func(__e *ControlFlow) {
V956 := __e.Get(1)
_ = V956
V957 := __e.Get(2)
_ = V957
tmp10034 := PrimEqual(Nil, V956)

if True == tmp10034 {
__e.Return(PrimCloseStream(V957))
return
} else {
tmp10032 := PrimIsPair(V956)

var ifres10028 Obj

if True == tmp10032 {
tmp10030 := PrimHead(V956)

tmp10031 := PrimIsPair(tmp10030)

var ifres10029 Obj

if True == tmp10031 {
ifres10029 = True


} else {
ifres10029 = False


}

ifres10028 = ifres10029


} else {
ifres10028 = False


}

if True == ifres10028 {
tmp10021 := PrimTail(V956)

tmp10022 := PrimHead(V956)

tmp10023 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp10022, V957)


_ = tmp10023

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp10021, V957)
return


} else {
tmp10026 := PrimIsPair(V956)

if True == tmp10026 {
tmp10024 := PrimTail(V956)

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp10024, V957)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.write-kl")))
return
}


}


}


}, 2)

tmp10035 := Call(__e, ns2_1set, symshen_4write_1kl, tmp10020)


_ = tmp10035

tmp10036 := MakeNative(func(__e *ControlFlow) {
V960 := __e.Get(1)
_ = V960
V961 := __e.Get(2)
_ = V961
tmp10076 := PrimIsPair(V960)

var ifres10039 Obj

if True == tmp10076 {
tmp10074 := PrimHead(V960)

tmp10075 := PrimEqual(symdefun, tmp10074)

var ifres10041 Obj

if True == tmp10075 {
tmp10072 := PrimTail(V960)

tmp10073 := PrimIsPair(tmp10072)

var ifres10043 Obj

if True == tmp10073 {
tmp10069 := PrimTail(V960)

tmp10070 := PrimHead(tmp10069)

tmp10071 := PrimEqual(symfail, tmp10070)

var ifres10045 Obj

if True == tmp10071 {
tmp10066 := PrimTail(V960)

tmp10067 := PrimTail(tmp10066)

tmp10068 := PrimIsPair(tmp10067)

var ifres10047 Obj

if True == tmp10068 {
tmp10062 := PrimTail(V960)

tmp10063 := PrimTail(tmp10062)

tmp10064 := PrimHead(tmp10063)

tmp10065 := PrimEqual(Nil, tmp10064)

var ifres10049 Obj

if True == tmp10065 {
tmp10058 := PrimTail(V960)

tmp10059 := PrimTail(tmp10058)

tmp10060 := PrimTail(tmp10059)

tmp10061 := PrimIsPair(tmp10060)

var ifres10051 Obj

if True == tmp10061 {
tmp10053 := PrimTail(V960)

tmp10054 := PrimTail(tmp10053)

tmp10055 := PrimTail(tmp10054)

tmp10056 := PrimTail(tmp10055)

tmp10057 := PrimEqual(Nil, tmp10056)

var ifres10052 Obj

if True == tmp10057 {
ifres10052 = True


} else {
ifres10052 = False


}

ifres10051 = ifres10052


} else {
ifres10051 = False


}

var ifres10050 Obj

if True == ifres10051 {
ifres10050 = True


} else {
ifres10050 = False


}

ifres10049 = ifres10050


} else {
ifres10049 = False


}

var ifres10048 Obj

if True == ifres10049 {
ifres10048 = True


} else {
ifres10048 = False


}

ifres10047 = ifres10048


} else {
ifres10047 = False


}

var ifres10046 Obj

if True == ifres10047 {
ifres10046 = True


} else {
ifres10046 = False


}

ifres10045 = ifres10046


} else {
ifres10045 = False


}

var ifres10044 Obj

if True == ifres10045 {
ifres10044 = True


} else {
ifres10044 = False


}

ifres10043 = ifres10044


} else {
ifres10043 = False


}

var ifres10042 Obj

if True == ifres10043 {
ifres10042 = True


} else {
ifres10042 = False


}

ifres10041 = ifres10042


} else {
ifres10041 = False


}

var ifres10040 Obj

if True == ifres10041 {
ifres10040 = True


} else {
ifres10040 = False


}

ifres10039 = ifres10040


} else {
ifres10039 = False


}

if True == ifres10039 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V961)
return
} else {
tmp10037 := Call(__e, PrimFunc(symshen_4app), V960, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp10037, V961)
return


}


}, 2)

tmp10077 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp10036)


_ = tmp10077

tmp10078 := MakeNative(func(__e *ControlFlow) {
V962 := __e.Get(1)
_ = V962
tmp10087 := PrimEqual(MakeString(""), V962)

if True == tmp10087 {
__e.Return(MakeString(".kl"))
return
} else {
tmp10085 := PrimEqual(MakeString(".shen"), V962)

if True == tmp10085 {
__e.Return(MakeString(".kl"))
return
} else {
tmp10083 := Call(__e, PrimFunc(symshen_4_7string_2), V962)


if True == tmp10083 {
tmp10079 := Call(__e, PrimFunc(symhdstr), V962)


tmp10080 := PrimTailString(V962)

tmp10081 := Call(__e, PrimFunc(symshen_4klfile), tmp10080)


__e.TailApply(PrimFunc(sym_8s), tmp10079, tmp10081)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.klfile")))
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp10078)
return




}, 0)

