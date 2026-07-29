package main

import . "github.com/pyrex41/shen-go/kl"

var LauncherMain = MakeNative(func(__e *ControlFlow) {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp19633 := MakeNative(func(__e *ControlFlow) {
V7102 := __e.Get(1)
_ = V7102
tmp19634 := MakeNative(func(__e *ControlFlow) {
W7103 := __e.Get(1)
_ = W7103
tmp19635 := MakeNative(func(__e *ControlFlow) {
Z7104 := __e.Get(1)
_ = Z7104
__e.TailApply(PrimFunc(symeval), Z7104)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp19635, W7103)
return


}, 1)

tmp19636 := Call(__e, PrimFunc(symread_1file), V7102)


__e.TailApply(tmp19634, tmp19636)
return


}, 1)

tmp19637 := Call(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp19633)


_ = tmp19637

tmp19638 := MakeNative(func(__e *ControlFlow) {
tmp19639 := Call(__e, PrimFunc(symversion))


tmp19640 := Call(__e, PrimFunc(symlanguage))


tmp19641 := Call(__e, PrimFunc(symport))


tmp19642 := PrimCons(tmp19641, Nil)

tmp19643 := PrimCons(tmp19640, tmp19642)

tmp19644 := Call(__e, PrimFunc(symimplementation))


tmp19645 := Call(__e, PrimFunc(symrelease))


tmp19646 := PrimCons(tmp19645, Nil)

tmp19647 := PrimCons(tmp19644, tmp19646)

tmp19648 := PrimCons(tmp19647, Nil)

tmp19649 := PrimCons(symimplementation, tmp19648)

tmp19650 := PrimCons(tmp19643, tmp19649)

tmp19651 := PrimCons(symport, tmp19650)

tmp19652 := Call(__e, PrimFunc(symshen_4app), tmp19651, MakeString("\n"), symshen_4r)


tmp19653 := PrimStringConcat(MakeString(" "), tmp19652)

__e.TailApply(PrimFunc(symshen_4app), tmp19639, tmp19653, symshen_4a)
return


}, 0)

tmp19654 := Call(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp19638)


_ = tmp19654

tmp19655 := MakeNative(func(__e *ControlFlow) {
V7105 := __e.Get(1)
_ = V7105
tmp19656 := Call(__e, PrimFunc(symshen_4app), V7105, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


__e.Return(PrimStringConcat(MakeString("Usage: "), tmp19656))
return


}, 1)

tmp19657 := Call(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp19655)


_ = tmp19657

tmp19658 := MakeNative(func(__e *ControlFlow) {
V7106 := __e.Get(1)
_ = V7106
tmp19665 := PrimEqual(Nil, V7106)

if True == tmp19665 {
__e.Return(PrimCons(symsuccess, Nil))
return
} else {
tmp19663 := PrimIsPair(V7106)

if True == tmp19663 {
tmp19659 := PrimHead(V7106)

tmp19660 := Call(__e, PrimFunc(symthaw), tmp19659)


_ = tmp19660

tmp19661 := PrimTail(V7106)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp19661)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
return
}


}


}, 1)

tmp19666 := Call(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp19658)


_ = tmp19666

tmp19667 := MakeNative(func(__e *ControlFlow) {
V7107 := __e.Get(1)
_ = V7107
tmp19668 := Call(__e, PrimFunc(symread_1from_1string), V7107)


tmp19669 := Call(__e, PrimFunc(symhead), tmp19668)


__e.TailApply(PrimFunc(symeval), tmp19669)
return


}, 1)

tmp19670 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp19667)


_ = tmp19670

tmp19671 := MakeNative(func(__e *ControlFlow) {
V7110 := __e.Get(1)
_ = V7110
tmp19681 := PrimEqual(MakeString("-e"), V7110)

if True == tmp19681 {
__e.Return(MakeString("--eval"))
return
} else {
tmp19679 := PrimEqual(MakeString("-l"), V7110)

if True == tmp19679 {
__e.Return(MakeString("--load"))
return
} else {
tmp19677 := PrimEqual(MakeString("-q"), V7110)

if True == tmp19677 {
__e.Return(MakeString("--quiet"))
return
} else {
tmp19675 := PrimEqual(MakeString("-s"), V7110)

if True == tmp19675 {
__e.Return(MakeString("--set"))
return
} else {
tmp19673 := PrimEqual(MakeString("-r"), V7110)

if True == tmp19673 {
__e.Return(MakeString("--repl"))
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

tmp19682 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp19671)


_ = tmp19682

tmp19683 := MakeNative(func(__e *ControlFlow) {
V7115 := __e.Get(1)
_ = V7115
V7116 := __e.Get(2)
_ = V7116
tmp19787 := PrimEqual(Nil, V7115)

if True == tmp19787 {
tmp19684 := Call(__e, PrimFunc(symreverse), V7116)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp19684)
return


} else {
tmp19785 := PrimIsPair(V7115)

var ifres19777 Obj

if True == tmp19785 {
tmp19783 := PrimHead(V7115)

tmp19784 := PrimEqual(MakeString("--eval"), tmp19783)

var ifres19779 Obj

if True == tmp19784 {
tmp19781 := PrimTail(V7115)

tmp19782 := PrimIsPair(tmp19781)

var ifres19780 Obj

if True == tmp19782 {
ifres19780 = True


} else {
ifres19780 = False


}

ifres19779 = ifres19780


} else {
ifres19779 = False


}

var ifres19778 Obj

if True == ifres19779 {
ifres19778 = True


} else {
ifres19778 = False


}

ifres19777 = ifres19778


} else {
ifres19777 = False


}

if True == ifres19777 {
tmp19685 := PrimTail(V7115)

tmp19686 := PrimTail(tmp19685)

tmp19687 := MakeNative(func(__e *ControlFlow) {
tmp19688 := PrimTail(V7115)

tmp19689 := PrimHead(tmp19688)

tmp19690 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp19689)


tmp19691 := Call(__e, PrimFunc(symshen_4app), tmp19690, MakeString("\n"), symshen_4a)


tmp19692 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp19691, tmp19692)
return


}, 0)

tmp19693 := PrimCons(tmp19687, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19686, tmp19693)
return


} else {
tmp19775 := PrimIsPair(V7115)

var ifres19767 Obj

if True == tmp19775 {
tmp19773 := PrimHead(V7115)

tmp19774 := PrimEqual(MakeString("--load"), tmp19773)

var ifres19769 Obj

if True == tmp19774 {
tmp19771 := PrimTail(V7115)

tmp19772 := PrimIsPair(tmp19771)

var ifres19770 Obj

if True == tmp19772 {
ifres19770 = True


} else {
ifres19770 = False


}

ifres19769 = ifres19770


} else {
ifres19769 = False


}

var ifres19768 Obj

if True == ifres19769 {
ifres19768 = True


} else {
ifres19768 = False


}

ifres19767 = ifres19768


} else {
ifres19767 = False


}

if True == ifres19767 {
tmp19694 := PrimTail(V7115)

tmp19695 := PrimTail(tmp19694)

tmp19696 := MakeNative(func(__e *ControlFlow) {
tmp19697 := PrimTail(V7115)

tmp19698 := PrimHead(tmp19697)

__e.TailApply(PrimFunc(symload), tmp19698)
return


}, 0)

tmp19699 := PrimCons(tmp19696, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19695, tmp19699)
return


} else {
tmp19765 := PrimIsPair(V7115)

var ifres19761 Obj

if True == tmp19765 {
tmp19763 := PrimHead(V7115)

tmp19764 := PrimEqual(MakeString("--quiet"), tmp19763)

var ifres19762 Obj

if True == tmp19764 {
ifres19762 = True


} else {
ifres19762 = False


}

ifres19761 = ifres19762


} else {
ifres19761 = False


}

if True == ifres19761 {
tmp19700 := PrimTail(V7115)

tmp19701 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSet(sym_dhush_d, True))
return
}, 0)

tmp19702 := PrimCons(tmp19701, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19700, tmp19702)
return


} else {
tmp19759 := PrimIsPair(V7115)

var ifres19746 Obj

if True == tmp19759 {
tmp19757 := PrimHead(V7115)

tmp19758 := PrimEqual(MakeString("--set"), tmp19757)

var ifres19748 Obj

if True == tmp19758 {
tmp19755 := PrimTail(V7115)

tmp19756 := PrimIsPair(tmp19755)

var ifres19750 Obj

if True == tmp19756 {
tmp19752 := PrimTail(V7115)

tmp19753 := PrimTail(tmp19752)

tmp19754 := PrimIsPair(tmp19753)

var ifres19751 Obj

if True == tmp19754 {
ifres19751 = True


} else {
ifres19751 = False


}

ifres19750 = ifres19751


} else {
ifres19750 = False


}

var ifres19749 Obj

if True == ifres19750 {
ifres19749 = True


} else {
ifres19749 = False


}

ifres19748 = ifres19749


} else {
ifres19748 = False


}

var ifres19747 Obj

if True == ifres19748 {
ifres19747 = True


} else {
ifres19747 = False


}

ifres19746 = ifres19747


} else {
ifres19746 = False


}

if True == ifres19746 {
tmp19703 := PrimTail(V7115)

tmp19704 := PrimTail(tmp19703)

tmp19705 := PrimTail(tmp19704)

tmp19706 := MakeNative(func(__e *ControlFlow) {
tmp19707 := PrimTail(V7115)

tmp19708 := PrimHead(tmp19707)

tmp19709 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp19708)


tmp19710 := PrimTail(V7115)

tmp19711 := PrimTail(tmp19710)

tmp19712 := PrimHead(tmp19711)

tmp19713 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp19712)


__e.Return(PrimSet(tmp19709, tmp19713))
return


}, 0)

tmp19714 := PrimCons(tmp19706, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19705, tmp19714)
return


} else {
tmp19744 := PrimIsPair(V7115)

var ifres19740 Obj

if True == tmp19744 {
tmp19742 := PrimHead(V7115)

tmp19743 := PrimEqual(MakeString("--repl"), tmp19742)

var ifres19741 Obj

if True == tmp19743 {
ifres19741 = True


} else {
ifres19741 = False


}

ifres19740 = ifres19741


} else {
ifres19740 = False


}

if True == ifres19740 {
tmp19715 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V7116)


_ = tmp19715

tmp19716 := PrimTail(V7115)

__e.Return(PrimCons(symlaunch_1repl, tmp19716))
return


} else {
tmp19717 := MakeNative(func(__e *ControlFlow) {
Freeze7119 := __e.Get(1)
_ = Freeze7119
tmp19731 := PrimIsPair(V7115)

if True == tmp19731 {
tmp19718 := MakeNative(func(__e *ControlFlow) {
Result7118 := __e.Get(1)
_ = Result7118
tmp19720 := Call(__e, PrimFunc(symfail))


tmp19721 := PrimEqual(Result7118, tmp19720)

if True == tmp19721 {
__e.TailApply(PrimFunc(symthaw), Freeze7119)
return
} else {
__e.Return(Result7118)
return
}


}, 1)

tmp19722 := MakeNative(func(__e *ControlFlow) {
W7117 := __e.Get(1)
_ = W7117
tmp19726 := PrimEqual(False, W7117)

if True == tmp19726 {
__e.TailApply(PrimFunc(symfail))
return
} else {
tmp19723 := PrimTail(V7115)

tmp19724 := PrimCons(W7117, tmp19723)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19724, V7116)
return


}


}, 1)

tmp19727 := PrimHead(V7115)

tmp19728 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1flag_1map), tmp19727)


tmp19729 := Call(__e, tmp19722, tmp19728)


__e.TailApply(tmp19718, tmp19729)
return


} else {
__e.TailApply(PrimFunc(symthaw), Freeze7119)
return
}


}, 1)

tmp19732 := MakeNative(func(__e *ControlFlow) {
tmp19738 := PrimIsPair(V7115)

if True == tmp19738 {
tmp19733 := PrimHead(V7115)

tmp19734 := Call(__e, PrimFunc(symshen_4app), tmp19733, MakeString(""), symshen_4a)


tmp19735 := PrimStringConcat(MakeString("Invalid eval argument: "), tmp19734)

tmp19736 := PrimCons(tmp19735, Nil)

__e.Return(PrimCons(symerror, tmp19736))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
return
}


}, 0)

__e.TailApply(tmp19717, tmp19732)
return


}


}


}


}


}


}


}, 2)

tmp19788 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp19683)


_ = tmp19788

tmp19789 := MakeNative(func(__e *ControlFlow) {
V7120 := __e.Get(1)
_ = V7120
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), V7120, Nil)
return
}, 1)

tmp19790 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp19789)


_ = tmp19790

tmp19791 := MakeNative(func(__e *ControlFlow) {
V7121 := __e.Get(1)
_ = V7121
V7122 := __e.Get(2)
_ = V7122
tmp19792 := PrimCons(V7121, V7122)

tmp19793 := PrimSet(sym_dargv_d, tmp19792)

_ = tmp19793

tmp19794 := Call(__e, PrimFunc(symshen_4x_4launcher_4quiet_1load), V7121)


_ = tmp19794

__e.Return(PrimCons(symsuccess, Nil))
return


}, 2)

tmp19795 := Call(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp19791)


_ = tmp19795

tmp19796 := MakeNative(func(__e *ControlFlow) {
V7123 := __e.Get(1)
_ = V7123
tmp19883 := PrimIsPair(V7123)

var ifres19879 Obj

if True == tmp19883 {
tmp19881 := PrimTail(V7123)

tmp19882 := PrimEqual(Nil, tmp19881)

var ifres19880 Obj

if True == tmp19882 {
ifres19880 = True


} else {
ifres19880 = False


}

ifres19879 = ifres19880


} else {
ifres19879 = False


}

if True == ifres19879 {
__e.Return(PrimCons(symlaunch_1repl, Nil))
return
} else {
tmp19877 := PrimIsPair(V7123)

var ifres19868 Obj

if True == tmp19877 {
tmp19875 := PrimTail(V7123)

tmp19876 := PrimIsPair(tmp19875)

var ifres19870 Obj

if True == tmp19876 {
tmp19872 := PrimTail(V7123)

tmp19873 := PrimHead(tmp19872)

tmp19874 := PrimEqual(MakeString("--help"), tmp19873)

var ifres19871 Obj

if True == tmp19874 {
ifres19871 = True


} else {
ifres19871 = False


}

ifres19870 = ifres19871


} else {
ifres19870 = False


}

var ifres19869 Obj

if True == ifres19870 {
ifres19869 = True


} else {
ifres19869 = False


}

ifres19868 = ifres19869


} else {
ifres19868 = False


}

if True == ifres19868 {
tmp19797 := PrimHead(V7123)

tmp19798 := Call(__e, PrimFunc(symshen_4x_4launcher_4help_1text), tmp19797)


tmp19799 := PrimCons(tmp19798, Nil)

__e.Return(PrimCons(symshow_1help, tmp19799))
return


} else {
tmp19866 := PrimIsPair(V7123)

var ifres19857 Obj

if True == tmp19866 {
tmp19864 := PrimTail(V7123)

tmp19865 := PrimIsPair(tmp19864)

var ifres19859 Obj

if True == tmp19865 {
tmp19861 := PrimTail(V7123)

tmp19862 := PrimHead(tmp19861)

tmp19863 := PrimEqual(MakeString("--version"), tmp19862)

var ifres19860 Obj

if True == tmp19863 {
ifres19860 = True


} else {
ifres19860 = False


}

ifres19859 = ifres19860


} else {
ifres19859 = False


}

var ifres19858 Obj

if True == ifres19859 {
ifres19858 = True


} else {
ifres19858 = False


}

ifres19857 = ifres19858


} else {
ifres19857 = False


}

if True == ifres19857 {
tmp19800 := Call(__e, PrimFunc(symshen_4x_4launcher_4version_1string))


tmp19801 := PrimCons(tmp19800, Nil)

__e.Return(PrimCons(symsuccess, tmp19801))
return


} else {
tmp19855 := PrimIsPair(V7123)

var ifres19846 Obj

if True == tmp19855 {
tmp19853 := PrimTail(V7123)

tmp19854 := PrimIsPair(tmp19853)

var ifres19848 Obj

if True == tmp19854 {
tmp19850 := PrimTail(V7123)

tmp19851 := PrimHead(tmp19850)

tmp19852 := PrimEqual(MakeString("repl"), tmp19851)

var ifres19849 Obj

if True == tmp19852 {
ifres19849 = True


} else {
ifres19849 = False


}

ifres19848 = ifres19849


} else {
ifres19848 = False


}

var ifres19847 Obj

if True == ifres19848 {
ifres19847 = True


} else {
ifres19847 = False


}

ifres19846 = ifres19847


} else {
ifres19846 = False


}

if True == ifres19846 {
tmp19802 := PrimTail(V7123)

tmp19803 := PrimTail(tmp19802)

__e.Return(PrimCons(symlaunch_1repl, tmp19803))
return


} else {
tmp19844 := PrimIsPair(V7123)

var ifres19830 Obj

if True == tmp19844 {
tmp19842 := PrimTail(V7123)

tmp19843 := PrimIsPair(tmp19842)

var ifres19832 Obj

if True == tmp19843 {
tmp19839 := PrimTail(V7123)

tmp19840 := PrimHead(tmp19839)

tmp19841 := PrimEqual(MakeString("script"), tmp19840)

var ifres19834 Obj

if True == tmp19841 {
tmp19836 := PrimTail(V7123)

tmp19837 := PrimTail(tmp19836)

tmp19838 := PrimIsPair(tmp19837)

var ifres19835 Obj

if True == tmp19838 {
ifres19835 = True


} else {
ifres19835 = False


}

ifres19834 = ifres19835


} else {
ifres19834 = False


}

var ifres19833 Obj

if True == ifres19834 {
ifres19833 = True


} else {
ifres19833 = False


}

ifres19832 = ifres19833


} else {
ifres19832 = False


}

var ifres19831 Obj

if True == ifres19832 {
ifres19831 = True


} else {
ifres19831 = False


}

ifres19830 = ifres19831


} else {
ifres19830 = False


}

if True == ifres19830 {
tmp19804 := PrimTail(V7123)

tmp19805 := PrimTail(tmp19804)

tmp19806 := PrimHead(tmp19805)

tmp19807 := PrimTail(V7123)

tmp19808 := PrimTail(tmp19807)

tmp19809 := PrimTail(tmp19808)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4script_1command), tmp19806, tmp19809)
return


} else {
tmp19828 := PrimIsPair(V7123)

var ifres19819 Obj

if True == tmp19828 {
tmp19826 := PrimTail(V7123)

tmp19827 := PrimIsPair(tmp19826)

var ifres19821 Obj

if True == tmp19827 {
tmp19823 := PrimTail(V7123)

tmp19824 := PrimHead(tmp19823)

tmp19825 := PrimEqual(MakeString("eval"), tmp19824)

var ifres19822 Obj

if True == tmp19825 {
ifres19822 = True


} else {
ifres19822 = False


}

ifres19821 = ifres19822


} else {
ifres19821 = False


}

var ifres19820 Obj

if True == ifres19821 {
ifres19820 = True


} else {
ifres19820 = False


}

ifres19819 = ifres19820


} else {
ifres19819 = False


}

if True == ifres19819 {
tmp19810 := PrimTail(V7123)

tmp19811 := PrimTail(tmp19810)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command), tmp19811)
return


} else {
tmp19817 := PrimIsPair(V7123)

var ifres19813 Obj

if True == tmp19817 {
tmp19815 := PrimTail(V7123)

tmp19816 := PrimIsPair(tmp19815)

var ifres19814 Obj

if True == tmp19816 {
ifres19814 = True


} else {
ifres19814 = False


}

ifres19813 = ifres19814


} else {
ifres19813 = False


}

if True == ifres19813 {
__e.Return(PrimCons(symunknown_1arguments, V7123))
return
} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4launch_1shen)
return
}


}


}


}


}


}


}


}, 1)

tmp19884 := Call(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp19796)


_ = tmp19884

tmp19885 := MakeNative(func(__e *ControlFlow) {
V7126 := __e.Get(1)
_ = V7126
tmp19984 := PrimIsPair(V7126)

var ifres19976 Obj

if True == tmp19984 {
tmp19982 := PrimHead(V7126)

tmp19983 := PrimEqual(symsuccess, tmp19982)

var ifres19978 Obj

if True == tmp19983 {
tmp19980 := PrimTail(V7126)

tmp19981 := PrimEqual(Nil, tmp19980)

var ifres19979 Obj

if True == tmp19981 {
ifres19979 = True


} else {
ifres19979 = False


}

ifres19978 = ifres19979


} else {
ifres19978 = False


}

var ifres19977 Obj

if True == ifres19978 {
ifres19977 = True


} else {
ifres19977 = False


}

ifres19976 = ifres19977


} else {
ifres19976 = False


}

if True == ifres19976 {
__e.Return(symshen_4x_4launcher_4done)
return
} else {
tmp19974 := PrimIsPair(V7126)

var ifres19961 Obj

if True == tmp19974 {
tmp19972 := PrimHead(V7126)

tmp19973 := PrimEqual(symsuccess, tmp19972)

var ifres19963 Obj

if True == tmp19973 {
tmp19970 := PrimTail(V7126)

tmp19971 := PrimIsPair(tmp19970)

var ifres19965 Obj

if True == tmp19971 {
tmp19967 := PrimTail(V7126)

tmp19968 := PrimTail(tmp19967)

tmp19969 := PrimEqual(Nil, tmp19968)

var ifres19966 Obj

if True == tmp19969 {
ifres19966 = True


} else {
ifres19966 = False


}

ifres19965 = ifres19966


} else {
ifres19965 = False


}

var ifres19964 Obj

if True == ifres19965 {
ifres19964 = True


} else {
ifres19964 = False


}

ifres19963 = ifres19964


} else {
ifres19963 = False


}

var ifres19962 Obj

if True == ifres19963 {
ifres19962 = True


} else {
ifres19962 = False


}

ifres19961 = ifres19962


} else {
ifres19961 = False


}

if True == ifres19961 {
tmp19886 := PrimTail(V7126)

tmp19887 := PrimHead(tmp19886)

tmp19888 := Call(__e, PrimFunc(symshen_4app), tmp19887, MakeString("\n"), symshen_4a)


tmp19889 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp19888, tmp19889)
return


} else {
tmp19959 := PrimIsPair(V7126)

var ifres19946 Obj

if True == tmp19959 {
tmp19957 := PrimHead(V7126)

tmp19958 := PrimEqual(symerror, tmp19957)

var ifres19948 Obj

if True == tmp19958 {
tmp19955 := PrimTail(V7126)

tmp19956 := PrimIsPair(tmp19955)

var ifres19950 Obj

if True == tmp19956 {
tmp19952 := PrimTail(V7126)

tmp19953 := PrimTail(tmp19952)

tmp19954 := PrimEqual(Nil, tmp19953)

var ifres19951 Obj

if True == tmp19954 {
ifres19951 = True


} else {
ifres19951 = False


}

ifres19950 = ifres19951


} else {
ifres19950 = False


}

var ifres19949 Obj

if True == ifres19950 {
ifres19949 = True


} else {
ifres19949 = False


}

ifres19948 = ifres19949


} else {
ifres19948 = False


}

var ifres19947 Obj

if True == ifres19948 {
ifres19947 = True


} else {
ifres19947 = False


}

ifres19946 = ifres19947


} else {
ifres19946 = False


}

if True == ifres19946 {
tmp19890 := PrimTail(V7126)

tmp19891 := PrimHead(tmp19890)

tmp19892 := Call(__e, PrimFunc(symshen_4app), tmp19891, MakeString("\n"), symshen_4a)


tmp19893 := PrimStringConcat(MakeString("ERROR: "), tmp19892)

tmp19894 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp19893, tmp19894)
return


} else {
tmp19944 := PrimIsPair(V7126)

var ifres19940 Obj

if True == tmp19944 {
tmp19942 := PrimHead(V7126)

tmp19943 := PrimEqual(symlaunch_1repl, tmp19942)

var ifres19941 Obj

if True == tmp19943 {
ifres19941 = True


} else {
ifres19941 = False


}

ifres19940 = ifres19941


} else {
ifres19940 = False


}

if True == ifres19940 {
__e.TailApply(PrimFunc(symshen_4repl))
return
} else {
tmp19938 := PrimIsPair(V7126)

var ifres19925 Obj

if True == tmp19938 {
tmp19936 := PrimHead(V7126)

tmp19937 := PrimEqual(symshow_1help, tmp19936)

var ifres19927 Obj

if True == tmp19937 {
tmp19934 := PrimTail(V7126)

tmp19935 := PrimIsPair(tmp19934)

var ifres19929 Obj

if True == tmp19935 {
tmp19931 := PrimTail(V7126)

tmp19932 := PrimTail(tmp19931)

tmp19933 := PrimEqual(Nil, tmp19932)

var ifres19930 Obj

if True == tmp19933 {
ifres19930 = True


} else {
ifres19930 = False


}

ifres19929 = ifres19930


} else {
ifres19929 = False


}

var ifres19928 Obj

if True == ifres19929 {
ifres19928 = True


} else {
ifres19928 = False


}

ifres19927 = ifres19928


} else {
ifres19927 = False


}

var ifres19926 Obj

if True == ifres19927 {
ifres19926 = True


} else {
ifres19926 = False


}

ifres19925 = ifres19926


} else {
ifres19925 = False


}

if True == ifres19925 {
tmp19895 := PrimTail(V7126)

tmp19896 := PrimHead(tmp19895)

tmp19897 := Call(__e, PrimFunc(symshen_4app), tmp19896, MakeString("\n"), symshen_4a)


tmp19898 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp19897, tmp19898)
return


} else {
tmp19923 := PrimIsPair(V7126)

var ifres19910 Obj

if True == tmp19923 {
tmp19921 := PrimHead(V7126)

tmp19922 := PrimEqual(symunknown_1arguments, tmp19921)

var ifres19912 Obj

if True == tmp19922 {
tmp19919 := PrimTail(V7126)

tmp19920 := PrimIsPair(tmp19919)

var ifres19914 Obj

if True == tmp19920 {
tmp19916 := PrimTail(V7126)

tmp19917 := PrimTail(tmp19916)

tmp19918 := PrimIsPair(tmp19917)

var ifres19915 Obj

if True == tmp19918 {
ifres19915 = True


} else {
ifres19915 = False


}

ifres19914 = ifres19915


} else {
ifres19914 = False


}

var ifres19913 Obj

if True == ifres19914 {
ifres19913 = True


} else {
ifres19913 = False


}

ifres19912 = ifres19913


} else {
ifres19912 = False


}

var ifres19911 Obj

if True == ifres19912 {
ifres19911 = True


} else {
ifres19911 = False


}

ifres19910 = ifres19911


} else {
ifres19910 = False


}

if True == ifres19910 {
tmp19899 := PrimTail(V7126)

tmp19900 := PrimTail(tmp19899)

tmp19901 := PrimHead(tmp19900)

tmp19902 := PrimTail(V7126)

tmp19903 := PrimHead(tmp19902)

tmp19904 := Call(__e, PrimFunc(symshen_4app), tmp19903, MakeString(" --help' for more information.\n"), symshen_4a)


tmp19905 := PrimStringConcat(MakeString("\nTry `"), tmp19904)

tmp19906 := Call(__e, PrimFunc(symshen_4app), tmp19901, tmp19905, symshen_4a)


tmp19907 := PrimStringConcat(MakeString("ERROR: Invalid argument: "), tmp19906)

tmp19908 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp19907, tmp19908)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4default_1handle_1result)
return
}


}


}


}


}


}


}, 1)

tmp19985 := Call(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp19885)


_ = tmp19985

tmp19986 := MakeNative(func(__e *ControlFlow) {
V7127 := __e.Get(1)
_ = V7127
tmp19987 := Call(__e, PrimFunc(symshen_4x_4launcher_4launch_1shen), V7127)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4default_1handle_1result), tmp19987)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4x_4launcher_4main, tmp19986)
return




}, 0)

var symeval_1kl = MakeSymbol("eval-kl")
var symY = MakeSymbol("Y")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symtrap_1error = MakeSymbol("trap-error")
var sym_5 = MakeSymbol("<")
var symis = MakeSymbol("is")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symfix = MakeSymbol("fix")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symeval = MakeSymbol("eval")
var symfst = MakeSymbol("fst")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symassoc = MakeSymbol("assoc")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symintern = MakeSymbol("intern")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symcond = MakeSymbol("cond")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symcd = MakeSymbol("cd")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symnumber_2 = MakeSymbol("number?")
var symretract = MakeSymbol("retract")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var sym_6_a = MakeSymbol(">=")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4repl = MakeSymbol("shen.repl")
var sym_8s = MakeSymbol("@s")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symstr = MakeSymbol("str")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4macro_1_8ch = MakeSymbol("shen.macro-@ch")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symunion = MakeSymbol("union")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symfn = MakeSymbol("fn")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symexternal = MakeSymbol("external")
var sym_dport_d = MakeSymbol("*port*")
var symrelease = MakeSymbol("release")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4tame = MakeSymbol("shen.tame")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symmapcan = MakeSymbol("mapcan")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symshen_4mod = MakeSymbol("shen.mod")
var sym_5_a = MakeSymbol("<=")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symFinish = MakeSymbol("Finish")
var symsubst = MakeSymbol("subst")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symL = MakeSymbol("L")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symput = MakeSymbol("put")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symatom_2 = MakeSymbol("atom?")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var symreturn = MakeSymbol("return")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symunprofile = MakeSymbol("unprofile")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symremove = MakeSymbol("remove")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4initialise_1lambda_1tables = MakeSymbol("shen.initialise-lambda-tables")
var symvalue = MakeSymbol("value")
var symnl = MakeSymbol("nl")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symStart = MakeSymbol("Start")
var symlanguage = MakeSymbol("language")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4consume = MakeSymbol("shen.consume")
var symshen_4prhush = MakeSymbol("shen.prhush")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4multiples = MakeSymbol("shen.multiples")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symstoutput = MakeSymbol("stoutput")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symZ = MakeSymbol("Z")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symRecord = MakeSymbol("Record")
var symnot = MakeSymbol("not")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symAction = MakeSymbol("Action")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symdo = MakeSymbol("do")
var symoptimise_2 = MakeSymbol("optimise?")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symelement_2 = MakeSymbol("element?")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symvector = MakeSymbol("vector")
var symunabsolute = MakeSymbol("unabsolute")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symport = MakeSymbol("port")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4reader_1error = MakeSymbol("shen.reader-error")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symmacroexpand = MakeSymbol("macroexpand")
var symcons = MakeSymbol("cons")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symread_1file = MakeSymbol("read-file")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symfreeze = MakeSymbol("freeze")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symsave = MakeSymbol("save")
var sym_e = MakeSymbol("&")
var symvector_1_6 = MakeSymbol("vector->")
var sym_a_a = MakeSymbol("==")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4rule_1_6body = MakeSymbol("shen.rule->body")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symdefine = MakeSymbol("define")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var sym__ = MakeSymbol("_")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var syminferences = MakeSymbol("inferences")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symstream = MakeSymbol("stream")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var syminput_7 = MakeSymbol("input+")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symlazy = MakeSymbol("lazy")
var symB = MakeSymbol("B")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symvariable_2 = MakeSymbol("variable?")
var symhdv = MakeSymbol("hdv")
var symshen_4remove_1bystanders = MakeSymbol("shen.remove-bystanders")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symS = MakeSymbol("S")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symbar_b = MakeSymbol("bar!")
var symclose = MakeSymbol("close")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symsqts = MakeSymbol("sqts")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symdeclare = MakeSymbol("declare")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symRemainder = MakeSymbol("Remainder")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symit = MakeSymbol("it")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4member = MakeSymbol("shen.member")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symNewV = MakeSymbol("NewV")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4_8c = MakeSymbol("shen.@c")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symfresh = MakeSymbol("fresh")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var sym_5_1address = MakeSymbol("<-address")
var symget_1time = MakeSymbol("get-time")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symcn = MakeSymbol("cn")
var symfunction = MakeSymbol("function")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symV = MakeSymbol("V")
var symhash = MakeSymbol("hash")
var symctxt = MakeSymbol("ctxt")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var syminternal = MakeSymbol("internal")
var symshen_4lch = MakeSymbol("shen.lch")
var symmap = MakeSymbol("map")
var symexception = MakeSymbol("exception")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symResult = MakeSymbol("Result")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4a = MakeSymbol("shen.a")
var symunput = MakeSymbol("unput")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symwhen = MakeSymbol("when")
var symspy = MakeSymbol("spy")
var symshow_1help = MakeSymbol("show-help")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symbound_2 = MakeSymbol("bound?")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symassertz = MakeSymbol("assertz")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symAssumptions = MakeSymbol("Assumptions")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symshen_4_8ch = MakeSymbol("shen.@ch")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4op2 = MakeSymbol("shen.op2")
var symfork = MakeSymbol("fork")
var symshen_4rdecons = MakeSymbol("shen.rdecons")
var symnull = MakeSymbol("null")
var symprotect = MakeSymbol("protect")
var sym_4_4_4 = MakeSymbol("...")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symfactorise_2 = MakeSymbol("factorise?")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var sym_c_4 = MakeSymbol("/.")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4f = MakeSymbol("shen.f")
var symforeign = MakeSymbol("foreign")
var symreverse = MakeSymbol("reverse")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4insert = MakeSymbol("shen.insert")
var symdefun = MakeSymbol("defun")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symW = MakeSymbol("W")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var sym_dargv_d = MakeSymbol("*argv*")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symlength = MakeSymbol("length")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4passive_1bind = MakeSymbol("shen.passive-bind")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symappend = MakeSymbol("append")
var symshen_4t = MakeSymbol("shen.t")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symhead = MakeSymbol("head")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symsystemf = MakeSymbol("systemf")
var symshen_4macro_1_8c = MakeSymbol("shen.macro-@c")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symshen_4r = MakeSymbol("shen.r")
var symdifference = MakeSymbol("difference")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symloaded = MakeSymbol("loaded")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symversion = MakeSymbol("version")
var symlet = MakeSymbol("let")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4partial_1parse_1failure_2 = MakeSymbol("shen.partial-parse-failure?")
var symempty_2 = MakeSymbol("empty?")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symuntrack = MakeSymbol("untrack")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symsuccess = MakeSymbol("success")
var sym_8v = MakeSymbol("@v")
var symprint = MakeSymbol("print")
var symsum = MakeSymbol("sum")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symstep = MakeSymbol("step")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symAssumption = MakeSymbol("Assumption")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symC = MakeSymbol("C")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symcons_2 = MakeSymbol("cons?")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var sym_5_1 = MakeSymbol("<-")
var symfactorise = MakeSymbol("factorise")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symopen = MakeSymbol("open")
var symfail = MakeSymbol("fail")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symor = MakeSymbol("or")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symcases = MakeSymbol("cases")
var symTime = MakeSymbol("Time")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symsymbol = MakeSymbol("symbol")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symtc = MakeSymbol("tc")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symlimit = MakeSymbol("limit")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symvector_2 = MakeSymbol("vector?")
var syminput = MakeSymbol("input")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symnumber = MakeSymbol("number")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4rule_1_6head = MakeSymbol("shen.rule->head")
var symshen_4consume_1clause = MakeSymbol("shen.consume-clause")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var sym_a = MakeSymbol("=")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symmode = MakeSymbol("mode")
var sym_1 = MakeSymbol("-")
var symtype = MakeSymbol("type")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symspy_2 = MakeSymbol("spy?")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symTl = MakeSymbol("Tl")
var symtlv = MakeSymbol("tlv")
var symshen_4beta = MakeSymbol("shen.beta")
var symtlstr = MakeSymbol("tlstr")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symcompile = MakeSymbol("compile")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symprolog_1memory = MakeSymbol("prolog-memory")
var sym_d = MakeSymbol("*")
var symvar_2 = MakeSymbol("var?")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symdatatypes = MakeSymbol("datatypes")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symabsolute = MakeSymbol("absolute")
var symshen_4g = MakeSymbol("shen.g")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symstring_2 = MakeSymbol("string?")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symps = MakeSymbol("ps")
var symin_1package = MakeSymbol("in-package")
var sym_1_6 = MakeSymbol("->")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symprolog_2 = MakeSymbol("prolog?")
var sym_6_6 = MakeSymbol(">>")
var symA = MakeSymbol("A")
var symtc_2 = MakeSymbol("tc?")
var symread = MakeSymbol("read")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen = MakeSymbol("shen")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symtuple_2 = MakeSymbol("tuple?")
var symimplementation = MakeSymbol("implementation")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var sympr = MakeSymbol("pr")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symasserta = MakeSymbol("asserta")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symadjoin = MakeSymbol("adjoin")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var sym_i = MakeSymbol("{")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var sympreclude = MakeSymbol("preclude")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var sym_a_a_6 = MakeSymbol("==>")
var sym_5_1vector = MakeSymbol("<-vector")
var symintersection = MakeSymbol("intersection")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symgensym = MakeSymbol("gensym")
var symexplode = MakeSymbol("explode")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symtrack = MakeSymbol("track")
var symlist = MakeSymbol("list")
var symshen_4change_1pointer_1value = MakeSymbol("shen.change-pointer-value")
var symprofile = MakeSymbol("profile")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symspecialise = MakeSymbol("specialise")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symboolean = MakeSymbol("boolean")
var symtracked = MakeSymbol("tracked")
var symconcat = MakeSymbol("concat")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var sym_dmacros_d = MakeSymbol("*macros*")
var symset = MakeSymbol("set")
var symshen_4goto = MakeSymbol("shen.goto")
var symHypotheses = MakeSymbol("Hypotheses")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symabsvector = MakeSymbol("absvector")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symwhere = MakeSymbol("where")
var symshen_4cons_1form_1respect_1modes = MakeSymbol("shen.cons-form-respect-modes")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symabsvector_2 = MakeSymbol("absvector?")
var symK = MakeSymbol("K")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symis_b = MakeSymbol("is!")
var symcall = MakeSymbol("call")
var symdefcc = MakeSymbol("defcc")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4passive_1variables = MakeSymbol("shen.passive-variables")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4side_1conditions_1_6goals = MakeSymbol("shen.side-conditions->goals")
var symuserdefs = MakeSymbol("userdefs")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4input_1h_7 = MakeSymbol("shen.input-h+")
var symshen_4premises_1_6goals = MakeSymbol("shen.premises->goals")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var sym_c = MakeSymbol("/")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4reader_1error_1message = MakeSymbol("shen.reader-error-message")
var symrun = MakeSymbol("run")
var symif = MakeSymbol("if")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symincluded = MakeSymbol("included")
var symreceive = MakeSymbol("receive")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4printF = MakeSymbol("shen.printF")
var sym_3 = MakeSymbol("$")
var symdatatype = MakeSymbol("datatype")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4correct = MakeSymbol("shen.correct")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symoutput = MakeSymbol("output")
var sym_e_e = MakeSymbol("&&")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symhush = MakeSymbol("hush")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symNewAssumptions = MakeSymbol("NewAssumptions")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symget = MakeSymbol("get")
var symtail = MakeSymbol("tail")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symParse = MakeSymbol("Parse")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symstep_2 = MakeSymbol("step?")
var sym_6 = MakeSymbol(">")
var symunix = MakeSymbol("unix")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var sym_5e_6 = MakeSymbol("<e>")
var symtl = MakeSymbol("tl")
var symshen_4macros = MakeSymbol("shen.macros")
var symbind = MakeSymbol("bind")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symlineread = MakeSymbol("lineread")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symFreeze = MakeSymbol("Freeze")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4process_1input_7 = MakeSymbol("shen.process-input+")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var sympackage_2 = MakeSymbol("package?")
var sym_dos_d = MakeSymbol("*os*")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4raise_1syntax_1error = MakeSymbol("shen.raise-syntax-error")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symporters = MakeSymbol("porters")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symunit = MakeSymbol("unit")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symKey = MakeSymbol("Key")
var symstring = MakeSymbol("string")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symabort = MakeSymbol("abort")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4specialise_1member = MakeSymbol("shen.specialise-member")
var symshen_4_1m = MakeSymbol("shen.-m")
var symoptimise = MakeSymbol("optimise")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symnth = MakeSymbol("nth")
var symhd = MakeSymbol("hd")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symthaw = MakeSymbol("thaw")
var symarity = MakeSymbol("arity")
var symin = MakeSymbol("in")
var syminclude = MakeSymbol("include")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symfindall = MakeSymbol("findall")
var symmake_1string = MakeSymbol("make-string")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symtime = MakeSymbol("time")
var symverified = MakeSymbol("verified")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symTm = MakeSymbol("Tm")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symSelect = MakeSymbol("Select")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symdestroy = MakeSymbol("destroy")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4syntax_1error_1message = MakeSymbol("shen.syntax-error-message")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4remove_1pointer = MakeSymbol("shen.remove-pointer")
var symsimple_1error = MakeSymbol("simple-error")
var sympos = MakeSymbol("pos")
var symMessage = MakeSymbol("Message")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symlambda = MakeSymbol("lambda")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symand = MakeSymbol("and")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4processed = MakeSymbol("shen.processed")
var symos = MakeSymbol("os")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4hush = MakeSymbol("shen.hush")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symX = MakeSymbol("X")
var symshen_4app = MakeSymbol("shen.app")
var symstinput = MakeSymbol("stinput")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symHd = MakeSymbol("Hd")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var sym_drelease_d = MakeSymbol("*release*")
var sympackage = MakeSymbol("package")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var sym_j = MakeSymbol("}")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4shen = MakeSymbol("shen.shen")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symfile = MakeSymbol("file")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4specialise_1consume = MakeSymbol("shen.specialise-consume")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var sym_8p = MakeSymbol("@p")
var symout = MakeSymbol("out")
var symload = MakeSymbol("load")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4cut = MakeSymbol("shen.cut")
var symGoTo = MakeSymbol("GoTo")
var symshen_4deref = MakeSymbol("shen.deref")
var sym_dabsolute_d = MakeSymbol("*absolute*")
var sym_5end_6 = MakeSymbol("<end>")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symsnd = MakeSymbol("snd")
var symhush_2 = MakeSymbol("hush?")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symu_b = MakeSymbol("u!")
var symP = MakeSymbol("P")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symhdstr = MakeSymbol("hdstr")
var sym_dlanguage_d = MakeSymbol("*language*")
var symerror = MakeSymbol("error")
var sym_b = MakeSymbol("!")
var symoccurs_2 = MakeSymbol("occurs?")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4member_1clause = MakeSymbol("shen.member-clause")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var sym_7 = MakeSymbol("+")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4_dlambdatable_d = MakeSymbol("shen.*lambdatable*")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
