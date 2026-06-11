package main

import . "github.com/tiancaiamao/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp1492 := MakeNative(func(__e *ControlFlow) {
V531 := __e.Get(1)
_ = V531
tmp1493 := MakeNative(func(__e *ControlFlow) {
W532 := __e.Get(1)
_ = W532
__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), W532)
return
}, 1)

tmp1494 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V531)


__e.TailApply(tmp1493, tmp1494)
return


}, 1)

tmp1495 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp1492)


_ = tmp1495

tmp1496 := MakeNative(func(__e *ControlFlow) {
V533 := __e.Get(1)
_ = V533
tmp1549 := PrimIsPair(V533)

var ifres1523 Obj

if True == tmp1549 {
tmp1547 := PrimHead(V533)

tmp1548 := PrimEqual(symdefun, tmp1547)

var ifres1525 Obj

if True == tmp1548 {
tmp1545 := PrimTail(V533)

tmp1546 := PrimIsPair(tmp1545)

var ifres1527 Obj

if True == tmp1546 {
tmp1542 := PrimTail(V533)

tmp1543 := PrimTail(tmp1542)

tmp1544 := PrimIsPair(tmp1543)

var ifres1529 Obj

if True == tmp1544 {
tmp1538 := PrimTail(V533)

tmp1539 := PrimTail(tmp1538)

tmp1540 := PrimTail(tmp1539)

tmp1541 := PrimIsPair(tmp1540)

var ifres1531 Obj

if True == tmp1541 {
tmp1533 := PrimTail(V533)

tmp1534 := PrimTail(tmp1533)

tmp1535 := PrimTail(tmp1534)

tmp1536 := PrimTail(tmp1535)

tmp1537 := PrimEqual(Nil, tmp1536)

var ifres1532 Obj

if True == tmp1537 {
ifres1532 = True


} else {
ifres1532 = False


}

ifres1531 = ifres1532


} else {
ifres1531 = False


}

var ifres1530 Obj

if True == ifres1531 {
ifres1530 = True


} else {
ifres1530 = False


}

ifres1529 = ifres1530


} else {
ifres1529 = False


}

var ifres1528 Obj

if True == ifres1529 {
ifres1528 = True


} else {
ifres1528 = False


}

ifres1527 = ifres1528


} else {
ifres1527 = False


}

var ifres1526 Obj

if True == ifres1527 {
ifres1526 = True


} else {
ifres1526 = False


}

ifres1525 = ifres1526


} else {
ifres1525 = False


}

var ifres1524 Obj

if True == ifres1525 {
ifres1524 = True


} else {
ifres1524 = False


}

ifres1523 = ifres1524


} else {
ifres1523 = False


}

if True == ifres1523 {
tmp1497 := MakeNative(func(__e *ControlFlow) {
W534 := __e.Get(1)
_ = W534
tmp1498 := MakeNative(func(__e *ControlFlow) {
W535 := __e.Get(1)
_ = W535
tmp1499 := MakeNative(func(__e *ControlFlow) {
W536 := __e.Get(1)
_ = W536
tmp1500 := MakeNative(func(__e *ControlFlow) {
W537 := __e.Get(1)
_ = W537
tmp1501 := PrimTail(V533)

tmp1502 := PrimHead(tmp1501)

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp1502)
return


}, 1)

tmp1503 := Call(__e, PrimFunc(symeval_1kl), V533)


__e.TailApply(tmp1500, tmp1503)
return


}, 1)

tmp1504 := PrimTail(V533)

tmp1505 := PrimHead(tmp1504)

tmp1506 := Call(__e, PrimFunc(symshen_4record_1kl), tmp1505, V533)


__e.TailApply(tmp1499, tmp1506)
return


}, 1)

tmp1507 := PrimTail(V533)

tmp1508 := PrimHead(tmp1507)

tmp1509 := PrimTail(V533)

tmp1510 := PrimTail(tmp1509)

tmp1511 := PrimHead(tmp1510)

tmp1512 := Call(__e, PrimFunc(symlength), tmp1511)


tmp1513 := Call(__e, PrimFunc(symshen_4store_1arity), tmp1508, tmp1512)


__e.TailApply(tmp1498, tmp1513)
return


}, 1)

tmp1519 := PrimTail(V533)

tmp1520 := PrimHead(tmp1519)

tmp1521 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp1520)


var ifres1514 Obj

if True == tmp1521 {
tmp1515 := PrimTail(V533)

tmp1516 := PrimHead(tmp1515)

tmp1517 := Call(__e, PrimFunc(symshen_4app), tmp1516, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp1518 := PrimSimpleError(tmp1517)

ifres1514 = tmp1518


} else {
ifres1514 = symshen_4skip


}

__e.TailApply(tmp1497, ifres1514)
return


} else {
__e.Return(V533)
return
}


}, 1)

tmp1550 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp1496)


_ = tmp1550

tmp1551 := MakeNative(func(__e *ControlFlow) {
V538 := __e.Get(1)
_ = V538
tmp1652 := PrimIsPair(V538)

var ifres1644 Obj

if True == tmp1652 {
tmp1650 := PrimHead(V538)

tmp1651 := PrimEqual(symdefine, tmp1650)

var ifres1646 Obj

if True == tmp1651 {
tmp1648 := PrimTail(V538)

tmp1649 := PrimIsPair(tmp1648)

var ifres1647 Obj

if True == tmp1649 {
ifres1647 = True


} else {
ifres1647 = False


}

ifres1646 = ifres1647


} else {
ifres1646 = False


}

var ifres1645 Obj

if True == ifres1646 {
ifres1645 = True


} else {
ifres1645 = False


}

ifres1644 = ifres1645


} else {
ifres1644 = False


}

if True == ifres1644 {
tmp1552 := PrimTail(V538)

tmp1553 := PrimHead(tmp1552)

tmp1554 := PrimTail(V538)

tmp1555 := PrimTail(tmp1554)

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp1553, tmp1555)
return


} else {
tmp1642 := PrimIsPair(V538)

var ifres1616 Obj

if True == tmp1642 {
tmp1640 := PrimHead(V538)

tmp1641 := PrimEqual(symdefun, tmp1640)

var ifres1618 Obj

if True == tmp1641 {
tmp1638 := PrimTail(V538)

tmp1639 := PrimIsPair(tmp1638)

var ifres1620 Obj

if True == tmp1639 {
tmp1635 := PrimTail(V538)

tmp1636 := PrimTail(tmp1635)

tmp1637 := PrimIsPair(tmp1636)

var ifres1622 Obj

if True == tmp1637 {
tmp1631 := PrimTail(V538)

tmp1632 := PrimTail(tmp1631)

tmp1633 := PrimTail(tmp1632)

tmp1634 := PrimIsPair(tmp1633)

var ifres1624 Obj

if True == tmp1634 {
tmp1626 := PrimTail(V538)

tmp1627 := PrimTail(tmp1626)

tmp1628 := PrimTail(tmp1627)

tmp1629 := PrimTail(tmp1628)

tmp1630 := PrimEqual(Nil, tmp1629)

var ifres1625 Obj

if True == tmp1630 {
ifres1625 = True


} else {
ifres1625 = False


}

ifres1624 = ifres1625


} else {
ifres1624 = False


}

var ifres1623 Obj

if True == ifres1624 {
ifres1623 = True


} else {
ifres1623 = False


}

ifres1622 = ifres1623


} else {
ifres1622 = False


}

var ifres1621 Obj

if True == ifres1622 {
ifres1621 = True


} else {
ifres1621 = False


}

ifres1620 = ifres1621


} else {
ifres1620 = False


}

var ifres1619 Obj

if True == ifres1620 {
ifres1619 = True


} else {
ifres1619 = False


}

ifres1618 = ifres1619


} else {
ifres1618 = False


}

var ifres1617 Obj

if True == ifres1618 {
ifres1617 = True


} else {
ifres1617 = False


}

ifres1616 = ifres1617


} else {
ifres1616 = False


}

if True == ifres1616 {
__e.Return(V538)
return
} else {
tmp1614 := PrimIsPair(V538)

var ifres1595 Obj

if True == tmp1614 {
tmp1612 := PrimHead(V538)

tmp1613 := PrimEqual(symtype, tmp1612)

var ifres1597 Obj

if True == tmp1613 {
tmp1610 := PrimTail(V538)

tmp1611 := PrimIsPair(tmp1610)

var ifres1599 Obj

if True == tmp1611 {
tmp1607 := PrimTail(V538)

tmp1608 := PrimTail(tmp1607)

tmp1609 := PrimIsPair(tmp1608)

var ifres1601 Obj

if True == tmp1609 {
tmp1603 := PrimTail(V538)

tmp1604 := PrimTail(tmp1603)

tmp1605 := PrimTail(tmp1604)

tmp1606 := PrimEqual(Nil, tmp1605)

var ifres1602 Obj

if True == tmp1606 {
ifres1602 = True


} else {
ifres1602 = False


}

ifres1601 = ifres1602


} else {
ifres1601 = False


}

var ifres1600 Obj

if True == ifres1601 {
ifres1600 = True


} else {
ifres1600 = False


}

ifres1599 = ifres1600


} else {
ifres1599 = False


}

var ifres1598 Obj

if True == ifres1599 {
ifres1598 = True


} else {
ifres1598 = False


}

ifres1597 = ifres1598


} else {
ifres1597 = False


}

var ifres1596 Obj

if True == ifres1597 {
ifres1596 = True


} else {
ifres1596 = False


}

ifres1595 = ifres1596


} else {
ifres1595 = False


}

if True == ifres1595 {
tmp1556 := PrimTail(V538)

tmp1557 := PrimHead(tmp1556)

tmp1558 := PrimTail(V538)

tmp1559 := PrimTail(tmp1558)

tmp1560 := PrimHead(tmp1559)

tmp1561 := Call(__e, PrimFunc(symshen_4rcons__form), tmp1560)


tmp1562 := PrimCons(tmp1561, Nil)

tmp1563 := PrimCons(tmp1557, tmp1562)

__e.Return(PrimCons(symtype, tmp1563))
return


} else {
tmp1593 := PrimIsPair(V538)

var ifres1574 Obj

if True == tmp1593 {
tmp1591 := PrimHead(V538)

tmp1592 := PrimEqual(syminput_7, tmp1591)

var ifres1576 Obj

if True == tmp1592 {
tmp1589 := PrimTail(V538)

tmp1590 := PrimIsPair(tmp1589)

var ifres1578 Obj

if True == tmp1590 {
tmp1586 := PrimTail(V538)

tmp1587 := PrimTail(tmp1586)

tmp1588 := PrimIsPair(tmp1587)

var ifres1580 Obj

if True == tmp1588 {
tmp1582 := PrimTail(V538)

tmp1583 := PrimTail(tmp1582)

tmp1584 := PrimTail(tmp1583)

tmp1585 := PrimEqual(Nil, tmp1584)

var ifres1581 Obj

if True == tmp1585 {
ifres1581 = True


} else {
ifres1581 = False


}

ifres1580 = ifres1581


} else {
ifres1580 = False


}

var ifres1579 Obj

if True == ifres1580 {
ifres1579 = True


} else {
ifres1579 = False


}

ifres1578 = ifres1579


} else {
ifres1578 = False


}

var ifres1577 Obj

if True == ifres1578 {
ifres1577 = True


} else {
ifres1577 = False


}

ifres1576 = ifres1577


} else {
ifres1576 = False


}

var ifres1575 Obj

if True == ifres1576 {
ifres1575 = True


} else {
ifres1575 = False


}

ifres1574 = ifres1575


} else {
ifres1574 = False


}

if True == ifres1574 {
tmp1564 := PrimTail(V538)

tmp1565 := PrimHead(tmp1564)

tmp1566 := Call(__e, PrimFunc(symshen_4rcons__form), tmp1565)


tmp1567 := PrimTail(V538)

tmp1568 := PrimTail(tmp1567)

tmp1569 := PrimCons(tmp1566, tmp1568)

__e.Return(PrimCons(syminput_7, tmp1569))
return


} else {
tmp1572 := PrimIsPair(V538)

if True == tmp1572 {
tmp1570 := MakeNative(func(__e *ControlFlow) {
Z539 := __e.Get(1)
_ = Z539
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z539)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp1570, V538)
return


} else {
__e.Return(V538)
return
}


}


}


}


}


}, 1)

tmp1653 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp1551)


_ = tmp1653

tmp1654 := MakeNative(func(__e *ControlFlow) {
V540 := __e.Get(1)
_ = V540
V541 := __e.Get(2)
_ = V541
tmp1655 := MakeNative(func(__e *ControlFlow) {
Z542 := __e.Get(1)
_ = Z542
__e.TailApply(PrimFunc(symshen_4_5define_6), Z542)
return
}, 1)

tmp1656 := PrimCons(V540, V541)

__e.TailApply(PrimFunc(symcompile), tmp1655, tmp1656)
return


}, 2)

tmp1657 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp1654)


_ = tmp1657

tmp1658 := MakeNative(func(__e *ControlFlow) {
V543 := __e.Get(1)
_ = V543
tmp1659 := MakeNative(func(__e *ControlFlow) {
W544 := __e.Get(1)
_ = W544
tmp1682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W544)


if True == tmp1682 {
tmp1660 := MakeNative(func(__e *ControlFlow) {
W555 := __e.Get(1)
_ = W555
tmp1662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W555)


if True == tmp1662 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W555)
return
}


}, 1)

tmp1663 := MakeNative(func(__e *ControlFlow) {
W556 := __e.Get(1)
_ = W556
tmp1678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W556)


if True == tmp1678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1664 := MakeNative(func(__e *ControlFlow) {
W557 := __e.Get(1)
_ = W557
tmp1665 := MakeNative(func(__e *ControlFlow) {
W558 := __e.Get(1)
_ = W558
tmp1666 := MakeNative(func(__e *ControlFlow) {
W559 := __e.Get(1)
_ = W559
tmp1673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W559)


if True == tmp1673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1667 := MakeNative(func(__e *ControlFlow) {
W560 := __e.Get(1)
_ = W560
tmp1668 := MakeNative(func(__e *ControlFlow) {
W561 := __e.Get(1)
_ = W561
tmp1669 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W557, W560)


__e.TailApply(PrimFunc(symshen_4comb), W561, tmp1669)
return


}, 1)

tmp1670 := Call(__e, PrimFunc(symshen_4in_1_6), W559)


__e.TailApply(tmp1668, tmp1670)
return


}, 1)

tmp1671 := Call(__e, PrimFunc(symshen_4_5_1out), W559)


__e.TailApply(tmp1667, tmp1671)
return


}


}, 1)

tmp1674 := Call(__e, PrimFunc(symshen_4_5rules_6), W558)


__e.TailApply(tmp1666, tmp1674)
return


}, 1)

tmp1675 := Call(__e, PrimFunc(symshen_4in_1_6), W556)


__e.TailApply(tmp1665, tmp1675)
return


}, 1)

tmp1676 := Call(__e, PrimFunc(symshen_4_5_1out), W556)


__e.TailApply(tmp1664, tmp1676)
return


}


}, 1)

tmp1679 := Call(__e, PrimFunc(symshen_4_5name_6), V543)


tmp1680 := Call(__e, tmp1663, tmp1679)


__e.TailApply(tmp1660, tmp1680)
return


} else {
__e.Return(W544)
return
}


}, 1)

tmp1683 := MakeNative(func(__e *ControlFlow) {
W545 := __e.Get(1)
_ = W545
tmp1712 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W545)


if True == tmp1712 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1684 := MakeNative(func(__e *ControlFlow) {
W546 := __e.Get(1)
_ = W546
tmp1685 := MakeNative(func(__e *ControlFlow) {
W547 := __e.Get(1)
_ = W547
tmp1708 := Call(__e, PrimFunc(symshen_4hds_a_2), W547, sym_i)


if True == tmp1708 {
tmp1686 := MakeNative(func(__e *ControlFlow) {
W548 := __e.Get(1)
_ = W548
tmp1687 := MakeNative(func(__e *ControlFlow) {
W549 := __e.Get(1)
_ = W549
tmp1704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W549)


if True == tmp1704 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1688 := MakeNative(func(__e *ControlFlow) {
W550 := __e.Get(1)
_ = W550
tmp1701 := Call(__e, PrimFunc(symshen_4hds_a_2), W550, sym_j)


if True == tmp1701 {
tmp1689 := MakeNative(func(__e *ControlFlow) {
W551 := __e.Get(1)
_ = W551
tmp1690 := MakeNative(func(__e *ControlFlow) {
W552 := __e.Get(1)
_ = W552
tmp1697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W552)


if True == tmp1697 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1691 := MakeNative(func(__e *ControlFlow) {
W553 := __e.Get(1)
_ = W553
tmp1692 := MakeNative(func(__e *ControlFlow) {
W554 := __e.Get(1)
_ = W554
tmp1693 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W546, W553)


__e.TailApply(PrimFunc(symshen_4comb), W554, tmp1693)
return


}, 1)

tmp1694 := Call(__e, PrimFunc(symshen_4in_1_6), W552)


__e.TailApply(tmp1692, tmp1694)
return


}, 1)

tmp1695 := Call(__e, PrimFunc(symshen_4_5_1out), W552)


__e.TailApply(tmp1691, tmp1695)
return


}


}, 1)

tmp1698 := Call(__e, PrimFunc(symshen_4_5rules_6), W551)


__e.TailApply(tmp1690, tmp1698)
return


}, 1)

tmp1699 := Call(__e, PrimFunc(symtail), W550)


__e.TailApply(tmp1689, tmp1699)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1702 := Call(__e, PrimFunc(symshen_4in_1_6), W549)


__e.TailApply(tmp1688, tmp1702)
return


}


}, 1)

tmp1705 := Call(__e, PrimFunc(symshen_4_5signature_6), W548)


__e.TailApply(tmp1687, tmp1705)
return


}, 1)

tmp1706 := Call(__e, PrimFunc(symtail), W547)


__e.TailApply(tmp1686, tmp1706)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1709 := Call(__e, PrimFunc(symshen_4in_1_6), W545)


__e.TailApply(tmp1685, tmp1709)
return


}, 1)

tmp1710 := Call(__e, PrimFunc(symshen_4_5_1out), W545)


__e.TailApply(tmp1684, tmp1710)
return


}


}, 1)

tmp1713 := Call(__e, PrimFunc(symshen_4_5name_6), V543)


tmp1714 := Call(__e, tmp1683, tmp1713)


__e.TailApply(tmp1659, tmp1714)
return


}, 1)

tmp1715 := Call(__e, ns2_1set, symshen_4_5define_6, tmp1658)


_ = tmp1715

tmp1716 := MakeNative(func(__e *ControlFlow) {
V562 := __e.Get(1)
_ = V562
V563 := __e.Get(2)
_ = V563
tmp1717 := MakeNative(func(__e *ControlFlow) {
W564 := __e.Get(1)
_ = W564
tmp1718 := MakeNative(func(__e *ControlFlow) {
W566 := __e.Get(1)
_ = W566
tmp1719 := MakeNative(func(__e *ControlFlow) {
W567 := __e.Get(1)
_ = W567
tmp1720 := MakeNative(func(__e *ControlFlow) {
W569 := __e.Get(1)
_ = W569
tmp1721 := MakeNative(func(__e *ControlFlow) {
W570 := __e.Get(1)
_ = W570
__e.Return(W570)
return
}, 1)

tmp1722 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V562, W569, W566)


tmp1723 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp1722)


__e.TailApply(tmp1721, tmp1723)
return


}, 1)

tmp1724 := Call(__e, PrimFunc(symshen_4unprotect), V563)


__e.TailApply(tmp1720, tmp1724)
return


}, 1)

tmp1725 := MakeNative(func(__e *ControlFlow) {
Z568 := __e.Get(1)
_ = Z568
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V562, Z568)
return
}, 1)

tmp1726 := Call(__e, PrimFunc(symmap), tmp1725, V563)


__e.TailApply(tmp1719, tmp1726)
return


}, 1)

tmp1727 := Call(__e, PrimFunc(symshen_4arity_1chk), V562, W564)


__e.TailApply(tmp1718, tmp1727)
return


}, 1)

tmp1728 := MakeNative(func(__e *ControlFlow) {
Z565 := __e.Get(1)
_ = Z565
__e.TailApply(PrimFunc(symfst), Z565)
return
}, 1)

tmp1729 := Call(__e, PrimFunc(symmap), tmp1728, V563)


__e.TailApply(tmp1717, tmp1729)
return


}, 2)

tmp1730 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp1716)


_ = tmp1730

tmp1731 := MakeNative(func(__e *ControlFlow) {
V571 := __e.Get(1)
_ = V571
tmp1757 := Call(__e, PrimFunc(symtuple_2), V571)


if True == tmp1757 {
tmp1732 := Call(__e, PrimFunc(symfst), V571)


tmp1733 := Call(__e, PrimFunc(symshen_4unprotect), tmp1732)


tmp1734 := Call(__e, PrimFunc(symsnd), V571)


tmp1735 := Call(__e, PrimFunc(symshen_4unprotect), tmp1734)


__e.TailApply(PrimFunc(sym_8p), tmp1733, tmp1735)
return


} else {
tmp1755 := PrimIsPair(V571)

var ifres1742 Obj

if True == tmp1755 {
tmp1753 := PrimHead(V571)

tmp1754 := PrimEqual(symprotect, tmp1753)

var ifres1744 Obj

if True == tmp1754 {
tmp1751 := PrimTail(V571)

tmp1752 := PrimIsPair(tmp1751)

var ifres1746 Obj

if True == tmp1752 {
tmp1748 := PrimTail(V571)

tmp1749 := PrimTail(tmp1748)

tmp1750 := PrimEqual(Nil, tmp1749)

var ifres1747 Obj

if True == tmp1750 {
ifres1747 = True


} else {
ifres1747 = False


}

ifres1746 = ifres1747


} else {
ifres1746 = False


}

var ifres1745 Obj

if True == ifres1746 {
ifres1745 = True


} else {
ifres1745 = False


}

ifres1744 = ifres1745


} else {
ifres1744 = False


}

var ifres1743 Obj

if True == ifres1744 {
ifres1743 = True


} else {
ifres1743 = False


}

ifres1742 = ifres1743


} else {
ifres1742 = False


}

if True == ifres1742 {
tmp1736 := PrimTail(V571)

tmp1737 := PrimHead(tmp1736)

__e.TailApply(PrimFunc(symshen_4unprotect), tmp1737)
return


} else {
tmp1740 := PrimIsPair(V571)

if True == tmp1740 {
tmp1738 := MakeNative(func(__e *ControlFlow) {
Z572 := __e.Get(1)
_ = Z572
__e.TailApply(PrimFunc(symshen_4unprotect), Z572)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp1738, V571)
return


} else {
__e.Return(V571)
return
}


}


}


}, 1)

tmp1758 := Call(__e, ns2_1set, symshen_4unprotect, tmp1731)


_ = tmp1758

tmp1759 := MakeNative(func(__e *ControlFlow) {
V573 := __e.Get(1)
_ = V573
tmp1760 := MakeNative(func(__e *ControlFlow) {
W574 := __e.Get(1)
_ = W574
tmp1762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W574)


if True == tmp1762 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W574)
return
}


}, 1)

tmp1778 := PrimIsPair(V573)

var ifres1763 Obj

if True == tmp1778 {
tmp1764 := MakeNative(func(__e *ControlFlow) {
W575 := __e.Get(1)
_ = W575
tmp1765 := MakeNative(func(__e *ControlFlow) {
W576 := __e.Get(1)
_ = W576
tmp1773 := PrimIsSymbol(W575)

var ifres1769 Obj

if True == tmp1773 {
tmp1771 := PrimIsVariable(W575)

tmp1772 := PrimNot(tmp1771)

var ifres1770 Obj

if True == tmp1772 {
ifres1770 = True


} else {
ifres1770 = False


}

ifres1769 = ifres1770


} else {
ifres1769 = False


}

var ifres1766 Obj

if True == ifres1769 {
ifres1766 = W575


} else {
tmp1767 := Call(__e, PrimFunc(symshen_4app), W575, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp1768 := PrimSimpleError(tmp1767)

ifres1766 = tmp1768


}

__e.TailApply(PrimFunc(symshen_4comb), W576, ifres1766)
return


}, 1)

tmp1774 := Call(__e, PrimFunc(symtail), V573)


__e.TailApply(tmp1765, tmp1774)
return


}, 1)

tmp1775 := Call(__e, PrimFunc(symhead), V573)


tmp1776 := Call(__e, tmp1764, tmp1775)


ifres1763 = tmp1776


} else {
tmp1777 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres1763 = tmp1777


}

__e.TailApply(tmp1760, ifres1763)
return


}, 1)

tmp1779 := Call(__e, ns2_1set, symshen_4_5name_6, tmp1759)


_ = tmp1779

tmp1780 := MakeNative(func(__e *ControlFlow) {
V577 := __e.Get(1)
_ = V577
tmp1781 := MakeNative(func(__e *ControlFlow) {
W578 := __e.Get(1)
_ = W578
tmp1793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W578)


if True == tmp1793 {
tmp1782 := MakeNative(func(__e *ControlFlow) {
W584 := __e.Get(1)
_ = W584
tmp1784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W584)


if True == tmp1784 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W584)
return
}


}, 1)

tmp1785 := MakeNative(func(__e *ControlFlow) {
W585 := __e.Get(1)
_ = W585
tmp1789 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W585)


if True == tmp1789 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1786 := MakeNative(func(__e *ControlFlow) {
W586 := __e.Get(1)
_ = W586
__e.TailApply(PrimFunc(symshen_4comb), W586, Nil)
return
}, 1)

tmp1787 := Call(__e, PrimFunc(symshen_4in_1_6), W585)


__e.TailApply(tmp1786, tmp1787)
return


}


}, 1)

tmp1790 := Call(__e, PrimFunc(sym_5e_6), V577)


tmp1791 := Call(__e, tmp1785, tmp1790)


__e.TailApply(tmp1782, tmp1791)
return


} else {
__e.Return(W578)
return
}


}, 1)

tmp1815 := PrimIsPair(V577)

var ifres1794 Obj

if True == tmp1815 {
tmp1795 := MakeNative(func(__e *ControlFlow) {
W579 := __e.Get(1)
_ = W579
tmp1796 := MakeNative(func(__e *ControlFlow) {
W580 := __e.Get(1)
_ = W580
tmp1797 := MakeNative(func(__e *ControlFlow) {
W581 := __e.Get(1)
_ = W581
tmp1809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W581)


if True == tmp1809 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1798 := MakeNative(func(__e *ControlFlow) {
W582 := __e.Get(1)
_ = W582
tmp1799 := MakeNative(func(__e *ControlFlow) {
W583 := __e.Get(1)
_ = W583
tmp1802 := PrimCons(sym_j, Nil)

tmp1803 := PrimCons(sym_i, tmp1802)

tmp1804 := Call(__e, PrimFunc(symelement_2), W579, tmp1803)


tmp1805 := PrimNot(tmp1804)

if True == tmp1805 {
tmp1800 := PrimCons(W579, W582)

__e.TailApply(PrimFunc(symshen_4comb), W583, tmp1800)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1806 := Call(__e, PrimFunc(symshen_4in_1_6), W581)


__e.TailApply(tmp1799, tmp1806)
return


}, 1)

tmp1807 := Call(__e, PrimFunc(symshen_4_5_1out), W581)


__e.TailApply(tmp1798, tmp1807)
return


}


}, 1)

tmp1810 := Call(__e, PrimFunc(symshen_4_5signature_6), W580)


__e.TailApply(tmp1797, tmp1810)
return


}, 1)

tmp1811 := Call(__e, PrimFunc(symtail), V577)


__e.TailApply(tmp1796, tmp1811)
return


}, 1)

tmp1812 := Call(__e, PrimFunc(symhead), V577)


tmp1813 := Call(__e, tmp1795, tmp1812)


ifres1794 = tmp1813


} else {
tmp1814 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres1794 = tmp1814


}

__e.TailApply(tmp1781, ifres1794)
return


}, 1)

tmp1816 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp1780)


_ = tmp1816

tmp1817 := MakeNative(func(__e *ControlFlow) {
V587 := __e.Get(1)
_ = V587
tmp1818 := MakeNative(func(__e *ControlFlow) {
W588 := __e.Get(1)
_ = W588
tmp1837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W588)


if True == tmp1837 {
tmp1819 := MakeNative(func(__e *ControlFlow) {
W595 := __e.Get(1)
_ = W595
tmp1821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W595)


if True == tmp1821 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W595)
return
}


}, 1)

tmp1822 := MakeNative(func(__e *ControlFlow) {
W596 := __e.Get(1)
_ = W596
tmp1833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W596)


if True == tmp1833 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1823 := MakeNative(func(__e *ControlFlow) {
W597 := __e.Get(1)
_ = W597
tmp1824 := MakeNative(func(__e *ControlFlow) {
W598 := __e.Get(1)
_ = W598
tmp1829 := Call(__e, PrimFunc(symempty_2), W597)


var ifres1825 Obj

if True == tmp1829 {
ifres1825 = Nil


} else {
tmp1826 := Call(__e, PrimFunc(symshen_4app), W597, MakeString("\n ..."), symshen_4r)


tmp1827 := PrimStringConcat(MakeString("Shen syntax error here:\n "), tmp1826)

tmp1828 := PrimSimpleError(tmp1827)

ifres1825 = tmp1828


}

__e.TailApply(PrimFunc(symshen_4comb), W598, ifres1825)
return


}, 1)

tmp1830 := Call(__e, PrimFunc(symshen_4in_1_6), W596)


__e.TailApply(tmp1824, tmp1830)
return


}, 1)

tmp1831 := Call(__e, PrimFunc(symshen_4_5_1out), W596)


__e.TailApply(tmp1823, tmp1831)
return


}


}, 1)

tmp1834 := Call(__e, PrimFunc(sym_5_b_6), V587)


tmp1835 := Call(__e, tmp1822, tmp1834)


__e.TailApply(tmp1819, tmp1835)
return


} else {
__e.Return(W588)
return
}


}, 1)

tmp1838 := MakeNative(func(__e *ControlFlow) {
W589 := __e.Get(1)
_ = W589
tmp1854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W589)


if True == tmp1854 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1839 := MakeNative(func(__e *ControlFlow) {
W590 := __e.Get(1)
_ = W590
tmp1840 := MakeNative(func(__e *ControlFlow) {
W591 := __e.Get(1)
_ = W591
tmp1841 := MakeNative(func(__e *ControlFlow) {
W592 := __e.Get(1)
_ = W592
tmp1849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W592)


if True == tmp1849 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1842 := MakeNative(func(__e *ControlFlow) {
W593 := __e.Get(1)
_ = W593
tmp1843 := MakeNative(func(__e *ControlFlow) {
W594 := __e.Get(1)
_ = W594
tmp1844 := Call(__e, PrimFunc(symshen_4linearise), W590)


tmp1845 := PrimCons(tmp1844, W593)

__e.TailApply(PrimFunc(symshen_4comb), W594, tmp1845)
return


}, 1)

tmp1846 := Call(__e, PrimFunc(symshen_4in_1_6), W592)


__e.TailApply(tmp1843, tmp1846)
return


}, 1)

tmp1847 := Call(__e, PrimFunc(symshen_4_5_1out), W592)


__e.TailApply(tmp1842, tmp1847)
return


}


}, 1)

tmp1850 := Call(__e, PrimFunc(symshen_4_5rules_6), W591)


__e.TailApply(tmp1841, tmp1850)
return


}, 1)

tmp1851 := Call(__e, PrimFunc(symshen_4in_1_6), W589)


__e.TailApply(tmp1840, tmp1851)
return


}, 1)

tmp1852 := Call(__e, PrimFunc(symshen_4_5_1out), W589)


__e.TailApply(tmp1839, tmp1852)
return


}


}, 1)

tmp1855 := Call(__e, PrimFunc(symshen_4_5rule_6), V587)


tmp1856 := Call(__e, tmp1838, tmp1855)


__e.TailApply(tmp1818, tmp1856)
return


}, 1)

tmp1857 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp1817)


_ = tmp1857

tmp1858 := MakeNative(func(__e *ControlFlow) {
V601 := __e.Get(1)
_ = V601
tmp1863 := Call(__e, PrimFunc(symtuple_2), V601)


if True == tmp1863 {
tmp1859 := Call(__e, PrimFunc(symfst), V601)


tmp1860 := Call(__e, PrimFunc(symfst), V601)


tmp1861 := Call(__e, PrimFunc(symsnd), V601)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1859, tmp1860, Nil, tmp1861)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise")))
return
}


}, 1)

tmp1864 := Call(__e, ns2_1set, symshen_4linearise, tmp1858)


_ = tmp1864

tmp1865 := MakeNative(func(__e *ControlFlow) {
V614 := __e.Get(1)
_ = V614
V615 := __e.Get(2)
_ = V615
V616 := __e.Get(3)
_ = V616
V617 := __e.Get(4)
_ = V617
tmp1903 := PrimEqual(Nil, V614)

if True == tmp1903 {
__e.TailApply(PrimFunc(sym_8p), V615, V617)
return
} else {
tmp1901 := PrimIsPair(V614)

var ifres1897 Obj

if True == tmp1901 {
tmp1899 := PrimHead(V614)

tmp1900 := PrimIsPair(tmp1899)

var ifres1898 Obj

if True == tmp1900 {
ifres1898 = True


} else {
ifres1898 = False


}

ifres1897 = ifres1898


} else {
ifres1897 = False


}

if True == ifres1897 {
tmp1866 := PrimHead(V614)

tmp1867 := PrimTail(V614)

tmp1868 := Call(__e, PrimFunc(symappend), tmp1866, tmp1867)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1868, V615, V616, V617)
return


} else {
tmp1895 := PrimIsPair(V614)

var ifres1891 Obj

if True == tmp1895 {
tmp1893 := PrimHead(V614)

tmp1894 := PrimIsVariable(tmp1893)

var ifres1892 Obj

if True == tmp1894 {
ifres1892 = True


} else {
ifres1892 = False


}

ifres1891 = ifres1892


} else {
ifres1891 = False


}

if True == ifres1891 {
tmp1885 := PrimHead(V614)

tmp1886 := Call(__e, PrimFunc(symelement_2), tmp1885, V616)


if True == tmp1886 {
tmp1869 := MakeNative(func(__e *ControlFlow) {
W618 := __e.Get(1)
_ = W618
tmp1870 := PrimTail(V614)

tmp1871 := PrimHead(V614)

tmp1872 := Call(__e, PrimFunc(symshen_4rep_1X), tmp1871, W618, V615)


tmp1873 := PrimHead(V614)

tmp1874 := PrimCons(tmp1873, Nil)

tmp1875 := PrimCons(W618, tmp1874)

tmp1876 := PrimCons(sym_a, tmp1875)

tmp1877 := PrimCons(V617, Nil)

tmp1878 := PrimCons(tmp1876, tmp1877)

tmp1879 := PrimCons(symwhere, tmp1878)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1870, tmp1872, V616, tmp1879)
return


}, 1)

tmp1880 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp1869, tmp1880)
return


} else {
tmp1881 := PrimTail(V614)

tmp1882 := PrimHead(V614)

tmp1883 := PrimCons(tmp1882, V616)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1881, V615, tmp1883, V617)
return


}


} else {
tmp1889 := PrimIsPair(V614)

if True == tmp1889 {
tmp1887 := PrimTail(V614)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp1887, V615, V616, V617)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise-h")))
return
}


}


}


}


}, 4)

tmp1904 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp1865)


_ = tmp1904

tmp1905 := MakeNative(func(__e *ControlFlow) {
V619 := __e.Get(1)
_ = V619
tmp1906 := MakeNative(func(__e *ControlFlow) {
W620 := __e.Get(1)
_ = W620
tmp1994 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W620)


if True == tmp1994 {
tmp1907 := MakeNative(func(__e *ControlFlow) {
W630 := __e.Get(1)
_ = W630
tmp1972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W630)


if True == tmp1972 {
tmp1908 := MakeNative(func(__e *ControlFlow) {
W637 := __e.Get(1)
_ = W637
tmp1935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W637)


if True == tmp1935 {
tmp1909 := MakeNative(func(__e *ControlFlow) {
W647 := __e.Get(1)
_ = W647
tmp1911 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W647)


if True == tmp1911 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W647)
return
}


}, 1)

tmp1912 := MakeNative(func(__e *ControlFlow) {
W648 := __e.Get(1)
_ = W648
tmp1931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W648)


if True == tmp1931 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1913 := MakeNative(func(__e *ControlFlow) {
W649 := __e.Get(1)
_ = W649
tmp1914 := MakeNative(func(__e *ControlFlow) {
W650 := __e.Get(1)
_ = W650
tmp1927 := Call(__e, PrimFunc(symshen_4hds_a_2), W650, sym_5_1)


if True == tmp1927 {
tmp1915 := MakeNative(func(__e *ControlFlow) {
W651 := __e.Get(1)
_ = W651
tmp1924 := PrimIsPair(W651)

if True == tmp1924 {
tmp1916 := MakeNative(func(__e *ControlFlow) {
W652 := __e.Get(1)
_ = W652
tmp1917 := MakeNative(func(__e *ControlFlow) {
W653 := __e.Get(1)
_ = W653
tmp1918 := PrimCons(W652, Nil)

tmp1919 := PrimCons(symshen_4choicepoint_b, tmp1918)

tmp1920 := Call(__e, PrimFunc(sym_8p), W649, tmp1919)


__e.TailApply(PrimFunc(symshen_4comb), W653, tmp1920)
return


}, 1)

tmp1921 := Call(__e, PrimFunc(symtail), W651)


__e.TailApply(tmp1917, tmp1921)
return


}, 1)

tmp1922 := Call(__e, PrimFunc(symhead), W651)


__e.TailApply(tmp1916, tmp1922)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1925 := Call(__e, PrimFunc(symtail), W650)


__e.TailApply(tmp1915, tmp1925)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1928 := Call(__e, PrimFunc(symshen_4in_1_6), W648)


__e.TailApply(tmp1914, tmp1928)
return


}, 1)

tmp1929 := Call(__e, PrimFunc(symshen_4_5_1out), W648)


__e.TailApply(tmp1913, tmp1929)
return


}


}, 1)

tmp1932 := Call(__e, PrimFunc(symshen_4_5patterns_6), V619)


tmp1933 := Call(__e, tmp1912, tmp1932)


__e.TailApply(tmp1909, tmp1933)
return


} else {
__e.Return(W637)
return
}


}, 1)

tmp1936 := MakeNative(func(__e *ControlFlow) {
W638 := __e.Get(1)
_ = W638
tmp1968 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W638)


if True == tmp1968 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1937 := MakeNative(func(__e *ControlFlow) {
W639 := __e.Get(1)
_ = W639
tmp1938 := MakeNative(func(__e *ControlFlow) {
W640 := __e.Get(1)
_ = W640
tmp1964 := Call(__e, PrimFunc(symshen_4hds_a_2), W640, sym_5_1)


if True == tmp1964 {
tmp1939 := MakeNative(func(__e *ControlFlow) {
W641 := __e.Get(1)
_ = W641
tmp1961 := PrimIsPair(W641)

if True == tmp1961 {
tmp1940 := MakeNative(func(__e *ControlFlow) {
W642 := __e.Get(1)
_ = W642
tmp1941 := MakeNative(func(__e *ControlFlow) {
W643 := __e.Get(1)
_ = W643
tmp1957 := Call(__e, PrimFunc(symshen_4hds_a_2), W643, symwhere)


if True == tmp1957 {
tmp1942 := MakeNative(func(__e *ControlFlow) {
W644 := __e.Get(1)
_ = W644
tmp1954 := PrimIsPair(W644)

if True == tmp1954 {
tmp1943 := MakeNative(func(__e *ControlFlow) {
W645 := __e.Get(1)
_ = W645
tmp1944 := MakeNative(func(__e *ControlFlow) {
W646 := __e.Get(1)
_ = W646
tmp1945 := PrimCons(W642, Nil)

tmp1946 := PrimCons(symshen_4choicepoint_b, tmp1945)

tmp1947 := PrimCons(tmp1946, Nil)

tmp1948 := PrimCons(W645, tmp1947)

tmp1949 := PrimCons(symwhere, tmp1948)

tmp1950 := Call(__e, PrimFunc(sym_8p), W639, tmp1949)


__e.TailApply(PrimFunc(symshen_4comb), W646, tmp1950)
return


}, 1)

tmp1951 := Call(__e, PrimFunc(symtail), W644)


__e.TailApply(tmp1944, tmp1951)
return


}, 1)

tmp1952 := Call(__e, PrimFunc(symhead), W644)


__e.TailApply(tmp1943, tmp1952)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1955 := Call(__e, PrimFunc(symtail), W643)


__e.TailApply(tmp1942, tmp1955)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1958 := Call(__e, PrimFunc(symtail), W641)


__e.TailApply(tmp1941, tmp1958)
return


}, 1)

tmp1959 := Call(__e, PrimFunc(symhead), W641)


__e.TailApply(tmp1940, tmp1959)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1962 := Call(__e, PrimFunc(symtail), W640)


__e.TailApply(tmp1939, tmp1962)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1965 := Call(__e, PrimFunc(symshen_4in_1_6), W638)


__e.TailApply(tmp1938, tmp1965)
return


}, 1)

tmp1966 := Call(__e, PrimFunc(symshen_4_5_1out), W638)


__e.TailApply(tmp1937, tmp1966)
return


}


}, 1)

tmp1969 := Call(__e, PrimFunc(symshen_4_5patterns_6), V619)


tmp1970 := Call(__e, tmp1936, tmp1969)


__e.TailApply(tmp1908, tmp1970)
return


} else {
__e.Return(W630)
return
}


}, 1)

tmp1973 := MakeNative(func(__e *ControlFlow) {
W631 := __e.Get(1)
_ = W631
tmp1990 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W631)


if True == tmp1990 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1974 := MakeNative(func(__e *ControlFlow) {
W632 := __e.Get(1)
_ = W632
tmp1975 := MakeNative(func(__e *ControlFlow) {
W633 := __e.Get(1)
_ = W633
tmp1986 := Call(__e, PrimFunc(symshen_4hds_a_2), W633, sym_1_6)


if True == tmp1986 {
tmp1976 := MakeNative(func(__e *ControlFlow) {
W634 := __e.Get(1)
_ = W634
tmp1983 := PrimIsPair(W634)

if True == tmp1983 {
tmp1977 := MakeNative(func(__e *ControlFlow) {
W635 := __e.Get(1)
_ = W635
tmp1978 := MakeNative(func(__e *ControlFlow) {
W636 := __e.Get(1)
_ = W636
tmp1979 := Call(__e, PrimFunc(sym_8p), W632, W635)


__e.TailApply(PrimFunc(symshen_4comb), W636, tmp1979)
return


}, 1)

tmp1980 := Call(__e, PrimFunc(symtail), W634)


__e.TailApply(tmp1978, tmp1980)
return


}, 1)

tmp1981 := Call(__e, PrimFunc(symhead), W634)


__e.TailApply(tmp1977, tmp1981)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1984 := Call(__e, PrimFunc(symtail), W633)


__e.TailApply(tmp1976, tmp1984)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp1987 := Call(__e, PrimFunc(symshen_4in_1_6), W631)


__e.TailApply(tmp1975, tmp1987)
return


}, 1)

tmp1988 := Call(__e, PrimFunc(symshen_4_5_1out), W631)


__e.TailApply(tmp1974, tmp1988)
return


}


}, 1)

tmp1991 := Call(__e, PrimFunc(symshen_4_5patterns_6), V619)


tmp1992 := Call(__e, tmp1973, tmp1991)


__e.TailApply(tmp1907, tmp1992)
return


} else {
__e.Return(W620)
return
}


}, 1)

tmp1995 := MakeNative(func(__e *ControlFlow) {
W621 := __e.Get(1)
_ = W621
tmp2025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W621)


if True == tmp2025 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp1996 := MakeNative(func(__e *ControlFlow) {
W622 := __e.Get(1)
_ = W622
tmp1997 := MakeNative(func(__e *ControlFlow) {
W623 := __e.Get(1)
_ = W623
tmp2021 := Call(__e, PrimFunc(symshen_4hds_a_2), W623, sym_1_6)


if True == tmp2021 {
tmp1998 := MakeNative(func(__e *ControlFlow) {
W624 := __e.Get(1)
_ = W624
tmp2018 := PrimIsPair(W624)

if True == tmp2018 {
tmp1999 := MakeNative(func(__e *ControlFlow) {
W625 := __e.Get(1)
_ = W625
tmp2000 := MakeNative(func(__e *ControlFlow) {
W626 := __e.Get(1)
_ = W626
tmp2014 := Call(__e, PrimFunc(symshen_4hds_a_2), W626, symwhere)


if True == tmp2014 {
tmp2001 := MakeNative(func(__e *ControlFlow) {
W627 := __e.Get(1)
_ = W627
tmp2011 := PrimIsPair(W627)

if True == tmp2011 {
tmp2002 := MakeNative(func(__e *ControlFlow) {
W628 := __e.Get(1)
_ = W628
tmp2003 := MakeNative(func(__e *ControlFlow) {
W629 := __e.Get(1)
_ = W629
tmp2004 := PrimCons(W625, Nil)

tmp2005 := PrimCons(W628, tmp2004)

tmp2006 := PrimCons(symwhere, tmp2005)

tmp2007 := Call(__e, PrimFunc(sym_8p), W622, tmp2006)


__e.TailApply(PrimFunc(symshen_4comb), W629, tmp2007)
return


}, 1)

tmp2008 := Call(__e, PrimFunc(symtail), W627)


__e.TailApply(tmp2003, tmp2008)
return


}, 1)

tmp2009 := Call(__e, PrimFunc(symhead), W627)


__e.TailApply(tmp2002, tmp2009)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2012 := Call(__e, PrimFunc(symtail), W626)


__e.TailApply(tmp2001, tmp2012)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2015 := Call(__e, PrimFunc(symtail), W624)


__e.TailApply(tmp2000, tmp2015)
return


}, 1)

tmp2016 := Call(__e, PrimFunc(symhead), W624)


__e.TailApply(tmp1999, tmp2016)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2019 := Call(__e, PrimFunc(symtail), W623)


__e.TailApply(tmp1998, tmp2019)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2022 := Call(__e, PrimFunc(symshen_4in_1_6), W621)


__e.TailApply(tmp1997, tmp2022)
return


}, 1)

tmp2023 := Call(__e, PrimFunc(symshen_4_5_1out), W621)


__e.TailApply(tmp1996, tmp2023)
return


}


}, 1)

tmp2026 := Call(__e, PrimFunc(symshen_4_5patterns_6), V619)


tmp2027 := Call(__e, tmp1995, tmp2026)


__e.TailApply(tmp1906, tmp2027)
return


}, 1)

tmp2028 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp1905)


_ = tmp2028

tmp2029 := MakeNative(func(__e *ControlFlow) {
V654 := __e.Get(1)
_ = V654
tmp2030 := MakeNative(func(__e *ControlFlow) {
W655 := __e.Get(1)
_ = W655
tmp2042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W655)


if True == tmp2042 {
tmp2031 := MakeNative(func(__e *ControlFlow) {
W662 := __e.Get(1)
_ = W662
tmp2033 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W662)


if True == tmp2033 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W662)
return
}


}, 1)

tmp2034 := MakeNative(func(__e *ControlFlow) {
W663 := __e.Get(1)
_ = W663
tmp2038 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W663)


if True == tmp2038 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2035 := MakeNative(func(__e *ControlFlow) {
W664 := __e.Get(1)
_ = W664
__e.TailApply(PrimFunc(symshen_4comb), W664, Nil)
return
}, 1)

tmp2036 := Call(__e, PrimFunc(symshen_4in_1_6), W663)


__e.TailApply(tmp2035, tmp2036)
return


}


}, 1)

tmp2039 := Call(__e, PrimFunc(sym_5e_6), V654)


tmp2040 := Call(__e, tmp2034, tmp2039)


__e.TailApply(tmp2031, tmp2040)
return


} else {
__e.Return(W655)
return
}


}, 1)

tmp2043 := MakeNative(func(__e *ControlFlow) {
W656 := __e.Get(1)
_ = W656
tmp2058 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W656)


if True == tmp2058 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2044 := MakeNative(func(__e *ControlFlow) {
W657 := __e.Get(1)
_ = W657
tmp2045 := MakeNative(func(__e *ControlFlow) {
W658 := __e.Get(1)
_ = W658
tmp2046 := MakeNative(func(__e *ControlFlow) {
W659 := __e.Get(1)
_ = W659
tmp2053 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W659)


if True == tmp2053 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2047 := MakeNative(func(__e *ControlFlow) {
W660 := __e.Get(1)
_ = W660
tmp2048 := MakeNative(func(__e *ControlFlow) {
W661 := __e.Get(1)
_ = W661
tmp2049 := PrimCons(W657, W660)

__e.TailApply(PrimFunc(symshen_4comb), W661, tmp2049)
return


}, 1)

tmp2050 := Call(__e, PrimFunc(symshen_4in_1_6), W659)


__e.TailApply(tmp2048, tmp2050)
return


}, 1)

tmp2051 := Call(__e, PrimFunc(symshen_4_5_1out), W659)


__e.TailApply(tmp2047, tmp2051)
return


}


}, 1)

tmp2054 := Call(__e, PrimFunc(symshen_4_5patterns_6), W658)


__e.TailApply(tmp2046, tmp2054)
return


}, 1)

tmp2055 := Call(__e, PrimFunc(symshen_4in_1_6), W656)


__e.TailApply(tmp2045, tmp2055)
return


}, 1)

tmp2056 := Call(__e, PrimFunc(symshen_4_5_1out), W656)


__e.TailApply(tmp2044, tmp2056)
return


}


}, 1)

tmp2059 := Call(__e, PrimFunc(symshen_4_5pattern_6), V654)


tmp2060 := Call(__e, tmp2043, tmp2059)


__e.TailApply(tmp2030, tmp2060)
return


}, 1)

tmp2061 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp2029)


_ = tmp2061

tmp2062 := MakeNative(func(__e *ControlFlow) {
V665 := __e.Get(1)
_ = V665
tmp2063 := MakeNative(func(__e *ControlFlow) {
W666 := __e.Get(1)
_ = W666
tmp2091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W666)


if True == tmp2091 {
tmp2064 := MakeNative(func(__e *ControlFlow) {
W673 := __e.Get(1)
_ = W673
tmp2078 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W673)


if True == tmp2078 {
tmp2065 := MakeNative(func(__e *ControlFlow) {
W676 := __e.Get(1)
_ = W676
tmp2067 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W676)


if True == tmp2067 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W676)
return
}


}, 1)

tmp2068 := MakeNative(func(__e *ControlFlow) {
W677 := __e.Get(1)
_ = W677
tmp2074 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W677)


if True == tmp2074 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2069 := MakeNative(func(__e *ControlFlow) {
W678 := __e.Get(1)
_ = W678
tmp2070 := MakeNative(func(__e *ControlFlow) {
W679 := __e.Get(1)
_ = W679
__e.TailApply(PrimFunc(symshen_4comb), W679, W678)
return
}, 1)

tmp2071 := Call(__e, PrimFunc(symshen_4in_1_6), W677)


__e.TailApply(tmp2070, tmp2071)
return


}, 1)

tmp2072 := Call(__e, PrimFunc(symshen_4_5_1out), W677)


__e.TailApply(tmp2069, tmp2072)
return


}


}, 1)

tmp2075 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V665)


tmp2076 := Call(__e, tmp2068, tmp2075)


__e.TailApply(tmp2065, tmp2076)
return


} else {
__e.Return(W673)
return
}


}, 1)

tmp2089 := PrimIsPair(V665)

var ifres2079 Obj

if True == tmp2089 {
tmp2080 := MakeNative(func(__e *ControlFlow) {
W674 := __e.Get(1)
_ = W674
tmp2081 := MakeNative(func(__e *ControlFlow) {
W675 := __e.Get(1)
_ = W675
tmp2084 := PrimIsPair(W674)

if True == tmp2084 {
tmp2082 := Call(__e, PrimFunc(symshen_4compound_1pattern), W674)


__e.TailApply(PrimFunc(symshen_4comb), W675, tmp2082)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2085 := Call(__e, PrimFunc(symtail), V665)


__e.TailApply(tmp2081, tmp2085)
return


}, 1)

tmp2086 := Call(__e, PrimFunc(symhead), V665)


tmp2087 := Call(__e, tmp2080, tmp2086)


ifres2079 = tmp2087


} else {
tmp2088 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2079 = tmp2088


}

__e.TailApply(tmp2064, ifres2079)
return


} else {
__e.Return(W666)
return
}


}, 1)

tmp2115 := Call(__e, PrimFunc(symshen_4ccons_2), V665)


var ifres2092 Obj

if True == tmp2115 {
tmp2093 := MakeNative(func(__e *ControlFlow) {
W667 := __e.Get(1)
_ = W667
tmp2094 := MakeNative(func(__e *ControlFlow) {
W668 := __e.Get(1)
_ = W668
tmp2110 := Call(__e, PrimFunc(symshen_4hds_a_2), W667, symvector)


if True == tmp2110 {
tmp2095 := MakeNative(func(__e *ControlFlow) {
W669 := __e.Get(1)
_ = W669
tmp2107 := Call(__e, PrimFunc(symshen_4hds_a_2), W669, MakeNumber(0))


if True == tmp2107 {
tmp2096 := MakeNative(func(__e *ControlFlow) {
W670 := __e.Get(1)
_ = W670
tmp2097 := MakeNative(func(__e *ControlFlow) {
W671 := __e.Get(1)
_ = W671
tmp2103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W671)


if True == tmp2103 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2098 := MakeNative(func(__e *ControlFlow) {
W672 := __e.Get(1)
_ = W672
tmp2099 := PrimCons(MakeNumber(0), Nil)

tmp2100 := PrimCons(symvector, tmp2099)

__e.TailApply(PrimFunc(symshen_4comb), W668, tmp2100)
return


}, 1)

tmp2101 := Call(__e, PrimFunc(symshen_4in_1_6), W671)


__e.TailApply(tmp2098, tmp2101)
return


}


}, 1)

tmp2104 := Call(__e, PrimFunc(sym_5end_6), W670)


__e.TailApply(tmp2097, tmp2104)
return


}, 1)

tmp2105 := Call(__e, PrimFunc(symtail), W669)


__e.TailApply(tmp2096, tmp2105)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2108 := Call(__e, PrimFunc(symtail), W667)


__e.TailApply(tmp2095, tmp2108)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2111 := Call(__e, PrimFunc(symtail), V665)


__e.TailApply(tmp2094, tmp2111)
return


}, 1)

tmp2112 := Call(__e, PrimFunc(symhead), V665)


tmp2113 := Call(__e, tmp2093, tmp2112)


ifres2092 = tmp2113


} else {
tmp2114 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2092 = tmp2114


}

__e.TailApply(tmp2063, ifres2092)
return


}, 1)

tmp2116 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp2062)


_ = tmp2116

tmp2117 := MakeNative(func(__e *ControlFlow) {
V680 := __e.Get(1)
_ = V680
tmp2118 := MakeNative(func(__e *ControlFlow) {
W681 := __e.Get(1)
_ = W681
tmp2120 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W681)


if True == tmp2120 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W681)
return
}


}, 1)

tmp2130 := PrimIsPair(V680)

var ifres2121 Obj

if True == tmp2130 {
tmp2122 := MakeNative(func(__e *ControlFlow) {
W682 := __e.Get(1)
_ = W682
tmp2123 := MakeNative(func(__e *ControlFlow) {
W683 := __e.Get(1)
_ = W683
tmp2125 := Call(__e, PrimFunc(symshen_4constructor_2), W682)


if True == tmp2125 {
__e.TailApply(PrimFunc(symshen_4comb), W683, W682)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2126 := Call(__e, PrimFunc(symtail), V680)


__e.TailApply(tmp2123, tmp2126)
return


}, 1)

tmp2127 := Call(__e, PrimFunc(symhead), V680)


tmp2128 := Call(__e, tmp2122, tmp2127)


ifres2121 = tmp2128


} else {
tmp2129 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2121 = tmp2129


}

__e.TailApply(tmp2118, ifres2121)
return


}, 1)

tmp2131 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp2117)


_ = tmp2131

tmp2132 := MakeNative(func(__e *ControlFlow) {
V684 := __e.Get(1)
_ = V684
tmp2133 := PrimCons(sym_8v, Nil)

tmp2134 := PrimCons(sym_8s, tmp2133)

tmp2135 := PrimCons(sym_8p, tmp2134)

tmp2136 := PrimCons(symcons, tmp2135)

__e.TailApply(PrimFunc(symelement_2), V684, tmp2136)
return


}, 1)

tmp2137 := Call(__e, ns2_1set, symshen_4constructor_2, tmp2132)


_ = tmp2137

tmp2138 := MakeNative(func(__e *ControlFlow) {
V685 := __e.Get(1)
_ = V685
tmp2139 := Call(__e, PrimFunc(symshen_4app), V685, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return(PrimSimpleError(tmp2139))
return


}, 1)

tmp2140 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp2138)


_ = tmp2140

tmp2141 := MakeNative(func(__e *ControlFlow) {
V686 := __e.Get(1)
_ = V686
V687 := __e.Get(2)
_ = V687
tmp2142 := MakeNative(func(__e *ControlFlow) {
W688 := __e.Get(1)
_ = W688
tmp2145 := PrimEqual(W688, False)

if True == tmp2145 {
__e.TailApply(PrimFunc(symthaw), V687)
return
} else {
tmp2143 := Call(__e, W688, V686)


__e.TailApply(tmp2143, V687)
return


}


}, 1)

tmp2146 := PrimValue(symshen_4_dcustom_1pattern_1compiler_d)

__e.TailApply(tmp2142, tmp2146)
return


}, 2)

tmp2147 := Call(__e, ns2_1set, symshen_4custom_1pattern_1compiler, tmp2141)


_ = tmp2147

tmp2148 := MakeNative(func(__e *ControlFlow) {
V689 := __e.Get(1)
_ = V689
tmp2149 := MakeNative(func(__e *ControlFlow) {
W690 := __e.Get(1)
_ = W690
tmp2151 := PrimEqual(W690, False)

if True == tmp2151 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.TailApply(W690, V689)
return
}


}, 1)

tmp2152 := PrimValue(symshen_4_dcustom_1pattern_1reducer_d)

__e.TailApply(tmp2149, tmp2152)
return


}, 1)

tmp2153 := Call(__e, ns2_1set, symshen_4custom_1pattern_1reducer, tmp2148)


_ = tmp2153

tmp2154 := MakeNative(func(__e *ControlFlow) {
V691 := __e.Get(1)
_ = V691
tmp2155 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4compound_1pattern_1h), V691)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4custom_1pattern_1compiler), V691, tmp2155)
return


}, 1)

tmp2156 := Call(__e, ns2_1set, symshen_4compound_1pattern, tmp2154)


_ = tmp2156

tmp2157 := MakeNative(func(__e *ControlFlow) {
V692 := __e.Get(1)
_ = V692
tmp2188 := PrimIsPair(V692)

var ifres2169 Obj

if True == tmp2188 {
tmp2186 := PrimTail(V692)

tmp2187 := PrimIsPair(tmp2186)

var ifres2171 Obj

if True == tmp2187 {
tmp2183 := PrimTail(V692)

tmp2184 := PrimTail(tmp2183)

tmp2185 := PrimIsPair(tmp2184)

var ifres2173 Obj

if True == tmp2185 {
tmp2179 := PrimTail(V692)

tmp2180 := PrimTail(tmp2179)

tmp2181 := PrimTail(tmp2180)

tmp2182 := PrimEqual(Nil, tmp2181)

var ifres2175 Obj

if True == tmp2182 {
tmp2177 := PrimHead(V692)

tmp2178 := Call(__e, PrimFunc(symshen_4constructor_2), tmp2177)


var ifres2176 Obj

if True == tmp2178 {
ifres2176 = True


} else {
ifres2176 = False


}

ifres2175 = ifres2176


} else {
ifres2175 = False


}

var ifres2174 Obj

if True == ifres2175 {
ifres2174 = True


} else {
ifres2174 = False


}

ifres2173 = ifres2174


} else {
ifres2173 = False


}

var ifres2172 Obj

if True == ifres2173 {
ifres2172 = True


} else {
ifres2172 = False


}

ifres2171 = ifres2172


} else {
ifres2171 = False


}

var ifres2170 Obj

if True == ifres2171 {
ifres2170 = True


} else {
ifres2170 = False


}

ifres2169 = ifres2170


} else {
ifres2169 = False


}

if True == ifres2169 {
tmp2158 := PrimHead(V692)

tmp2159 := PrimTail(V692)

tmp2160 := PrimHead(tmp2159)

tmp2161 := Call(__e, PrimFunc(symshen_4compile_1pattern_1fragment), tmp2160)


tmp2162 := PrimTail(V692)

tmp2163 := PrimTail(tmp2162)

tmp2164 := PrimHead(tmp2163)

tmp2165 := Call(__e, PrimFunc(symshen_4compile_1pattern_1fragment), tmp2164)


tmp2166 := PrimCons(tmp2165, Nil)

tmp2167 := PrimCons(tmp2161, tmp2166)

__e.Return(PrimCons(tmp2158, tmp2167))
return


} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V692)
return
}


}, 1)

tmp2189 := Call(__e, ns2_1set, symshen_4compound_1pattern_1h, tmp2157)


_ = tmp2189

tmp2190 := MakeNative(func(__e *ControlFlow) {
V693 := __e.Get(1)
_ = V693
tmp2219 := PrimIsPair(V693)

var ifres2201 Obj

if True == tmp2219 {
tmp2217 := PrimHead(V693)

tmp2218 := PrimEqual(symvector, tmp2217)

var ifres2203 Obj

if True == tmp2218 {
tmp2215 := PrimTail(V693)

tmp2216 := PrimIsPair(tmp2215)

var ifres2205 Obj

if True == tmp2216 {
tmp2212 := PrimTail(V693)

tmp2213 := PrimHead(tmp2212)

tmp2214 := PrimEqual(MakeNumber(0), tmp2213)

var ifres2207 Obj

if True == tmp2214 {
tmp2209 := PrimTail(V693)

tmp2210 := PrimTail(tmp2209)

tmp2211 := PrimEqual(Nil, tmp2210)

var ifres2208 Obj

if True == tmp2211 {
ifres2208 = True


} else {
ifres2208 = False


}

ifres2207 = ifres2208


} else {
ifres2207 = False


}

var ifres2206 Obj

if True == ifres2207 {
ifres2206 = True


} else {
ifres2206 = False


}

ifres2205 = ifres2206


} else {
ifres2205 = False


}

var ifres2204 Obj

if True == ifres2205 {
ifres2204 = True


} else {
ifres2204 = False


}

ifres2203 = ifres2204


} else {
ifres2203 = False


}

var ifres2202 Obj

if True == ifres2203 {
ifres2202 = True


} else {
ifres2202 = False


}

ifres2201 = ifres2202


} else {
ifres2201 = False


}

if True == ifres2201 {
__e.Return(V693)
return
} else {
tmp2199 := PrimIsPair(V693)

if True == tmp2199 {
__e.TailApply(PrimFunc(symshen_4compound_1pattern), V693)
return
} else {
tmp2197 := PrimEqual(V693, sym__)

if True == tmp2197 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp2192 := PrimCons(sym_5_1, Nil)

tmp2193 := PrimCons(sym_1_6, tmp2192)

tmp2194 := Call(__e, PrimFunc(symelement_2), V693, tmp2193)


tmp2195 := PrimNot(tmp2194)

if True == tmp2195 {
__e.Return(V693)
return
} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V693)
return
}


}


}


}


}, 1)

tmp2220 := Call(__e, ns2_1set, symshen_4compile_1pattern_1fragment, tmp2190)


_ = tmp2220

tmp2221 := MakeNative(func(__e *ControlFlow) {
V698 := __e.Get(1)
_ = V698
tmp2247 := PrimIsPair(V698)

var ifres2223 Obj

if True == tmp2247 {
tmp2245 := PrimHead(V698)

tmp2246 := PrimEqual(sym_8p, tmp2245)

var ifres2225 Obj

if True == tmp2246 {
tmp2243 := PrimTail(V698)

tmp2244 := PrimIsPair(tmp2243)

var ifres2227 Obj

if True == tmp2244 {
tmp2240 := PrimTail(V698)

tmp2241 := PrimHead(tmp2240)

tmp2242 := PrimEqual(symshen_4custom_1pattern, tmp2241)

var ifres2229 Obj

if True == tmp2242 {
tmp2237 := PrimTail(V698)

tmp2238 := PrimTail(tmp2237)

tmp2239 := PrimIsPair(tmp2238)

var ifres2231 Obj

if True == tmp2239 {
tmp2233 := PrimTail(V698)

tmp2234 := PrimTail(tmp2233)

tmp2235 := PrimTail(tmp2234)

tmp2236 := PrimEqual(Nil, tmp2235)

var ifres2232 Obj

if True == tmp2236 {
ifres2232 = True


} else {
ifres2232 = False


}

ifres2231 = ifres2232


} else {
ifres2231 = False


}

var ifres2230 Obj

if True == ifres2231 {
ifres2230 = True


} else {
ifres2230 = False


}

ifres2229 = ifres2230


} else {
ifres2229 = False


}

var ifres2228 Obj

if True == ifres2229 {
ifres2228 = True


} else {
ifres2228 = False


}

ifres2227 = ifres2228


} else {
ifres2227 = False


}

var ifres2226 Obj

if True == ifres2227 {
ifres2226 = True


} else {
ifres2226 = False


}

ifres2225 = ifres2226


} else {
ifres2225 = False


}

var ifres2224 Obj

if True == ifres2225 {
ifres2224 = True


} else {
ifres2224 = False


}

ifres2223 = ifres2224


} else {
ifres2223 = False


}

if True == ifres2223 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp2248 := Call(__e, ns2_1set, symshen_4custom_1pattern_2, tmp2221)


_ = tmp2248

tmp2249 := MakeNative(func(__e *ControlFlow) {
V701 := __e.Get(1)
_ = V701
tmp2277 := PrimIsPair(V701)

var ifres2253 Obj

if True == tmp2277 {
tmp2275 := PrimHead(V701)

tmp2276 := PrimEqual(sym_8p, tmp2275)

var ifres2255 Obj

if True == tmp2276 {
tmp2273 := PrimTail(V701)

tmp2274 := PrimIsPair(tmp2273)

var ifres2257 Obj

if True == tmp2274 {
tmp2270 := PrimTail(V701)

tmp2271 := PrimHead(tmp2270)

tmp2272 := PrimEqual(symshen_4custom_1pattern, tmp2271)

var ifres2259 Obj

if True == tmp2272 {
tmp2267 := PrimTail(V701)

tmp2268 := PrimTail(tmp2267)

tmp2269 := PrimIsPair(tmp2268)

var ifres2261 Obj

if True == tmp2269 {
tmp2263 := PrimTail(V701)

tmp2264 := PrimTail(tmp2263)

tmp2265 := PrimTail(tmp2264)

tmp2266 := PrimEqual(Nil, tmp2265)

var ifres2262 Obj

if True == tmp2266 {
ifres2262 = True


} else {
ifres2262 = False


}

ifres2261 = ifres2262


} else {
ifres2261 = False


}

var ifres2260 Obj

if True == ifres2261 {
ifres2260 = True


} else {
ifres2260 = False


}

ifres2259 = ifres2260


} else {
ifres2259 = False


}

var ifres2258 Obj

if True == ifres2259 {
ifres2258 = True


} else {
ifres2258 = False


}

ifres2257 = ifres2258


} else {
ifres2257 = False


}

var ifres2256 Obj

if True == ifres2257 {
ifres2256 = True


} else {
ifres2256 = False


}

ifres2255 = ifres2256


} else {
ifres2255 = False


}

var ifres2254 Obj

if True == ifres2255 {
ifres2254 = True


} else {
ifres2254 = False


}

ifres2253 = ifres2254


} else {
ifres2253 = False


}

if True == ifres2253 {
tmp2250 := PrimTail(V701)

tmp2251 := PrimTail(tmp2250)

__e.Return(PrimHead(tmp2251))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.custom-pattern-body")))
return
}


}, 1)

tmp2278 := Call(__e, ns2_1set, symshen_4custom_1pattern_1body, tmp2249)


_ = tmp2278

tmp2279 := MakeNative(func(__e *ControlFlow) {
V702 := __e.Get(1)
_ = V702
tmp2280 := MakeNative(func(__e *ControlFlow) {
W703 := __e.Get(1)
_ = W703
tmp2298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W703)


if True == tmp2298 {
tmp2281 := MakeNative(func(__e *ControlFlow) {
W706 := __e.Get(1)
_ = W706
tmp2283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W706)


if True == tmp2283 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W706)
return
}


}, 1)

tmp2296 := PrimIsPair(V702)

var ifres2284 Obj

if True == tmp2296 {
tmp2285 := MakeNative(func(__e *ControlFlow) {
W707 := __e.Get(1)
_ = W707
tmp2286 := MakeNative(func(__e *ControlFlow) {
W708 := __e.Get(1)
_ = W708
tmp2288 := PrimCons(sym_5_1, Nil)

tmp2289 := PrimCons(sym_1_6, tmp2288)

tmp2290 := Call(__e, PrimFunc(symelement_2), W707, tmp2289)


tmp2291 := PrimNot(tmp2290)

if True == tmp2291 {
__e.TailApply(PrimFunc(symshen_4comb), W708, W707)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2292 := Call(__e, PrimFunc(symtail), V702)


__e.TailApply(tmp2286, tmp2292)
return


}, 1)

tmp2293 := Call(__e, PrimFunc(symhead), V702)


tmp2294 := Call(__e, tmp2285, tmp2293)


ifres2284 = tmp2294


} else {
tmp2295 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2284 = tmp2295


}

__e.TailApply(tmp2281, ifres2284)
return


} else {
__e.Return(W703)
return
}


}, 1)

tmp2309 := PrimIsPair(V702)

var ifres2299 Obj

if True == tmp2309 {
tmp2300 := MakeNative(func(__e *ControlFlow) {
W704 := __e.Get(1)
_ = W704
tmp2301 := MakeNative(func(__e *ControlFlow) {
W705 := __e.Get(1)
_ = W705
tmp2304 := PrimEqual(W704, sym__)

if True == tmp2304 {
tmp2302 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W705, tmp2302)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2305 := Call(__e, PrimFunc(symtail), V702)


__e.TailApply(tmp2301, tmp2305)
return


}, 1)

tmp2306 := Call(__e, PrimFunc(symhead), V702)


tmp2307 := Call(__e, tmp2300, tmp2306)


ifres2299 = tmp2307


} else {
tmp2308 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2299 = tmp2308


}

__e.TailApply(tmp2280, ifres2299)
return


}, 1)

tmp2310 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp2279)


_ = tmp2310

tmp2311 := MakeNative(func(__e *ControlFlow) {
V709 := __e.Get(1)
_ = V709
tmp2312 := MakeNative(func(__e *ControlFlow) {
W710 := __e.Get(1)
_ = W710
tmp2314 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W710)


if True == tmp2314 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W710)
return
}


}, 1)

tmp2315 := MakeNative(func(__e *ControlFlow) {
W711 := __e.Get(1)
_ = W711
tmp2321 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W711)


if True == tmp2321 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2316 := MakeNative(func(__e *ControlFlow) {
W712 := __e.Get(1)
_ = W712
tmp2317 := MakeNative(func(__e *ControlFlow) {
W713 := __e.Get(1)
_ = W713
__e.TailApply(PrimFunc(symshen_4comb), W713, W712)
return
}, 1)

tmp2318 := Call(__e, PrimFunc(symshen_4in_1_6), W711)


__e.TailApply(tmp2317, tmp2318)
return


}, 1)

tmp2319 := Call(__e, PrimFunc(symshen_4_5_1out), W711)


__e.TailApply(tmp2316, tmp2319)
return


}


}, 1)

tmp2322 := Call(__e, PrimFunc(symshen_4_5pattern_6), V709)


tmp2323 := Call(__e, tmp2315, tmp2322)


__e.TailApply(tmp2312, tmp2323)
return


}, 1)

tmp2324 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp2311)


_ = tmp2324

tmp2325 := MakeNative(func(__e *ControlFlow) {
V714 := __e.Get(1)
_ = V714
tmp2326 := MakeNative(func(__e *ControlFlow) {
W715 := __e.Get(1)
_ = W715
tmp2328 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W715)


if True == tmp2328 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W715)
return
}


}, 1)

tmp2329 := MakeNative(func(__e *ControlFlow) {
W716 := __e.Get(1)
_ = W716
tmp2335 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W716)


if True == tmp2335 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2330 := MakeNative(func(__e *ControlFlow) {
W717 := __e.Get(1)
_ = W717
tmp2331 := MakeNative(func(__e *ControlFlow) {
W718 := __e.Get(1)
_ = W718
__e.TailApply(PrimFunc(symshen_4comb), W718, W717)
return
}, 1)

tmp2332 := Call(__e, PrimFunc(symshen_4in_1_6), W716)


__e.TailApply(tmp2331, tmp2332)
return


}, 1)

tmp2333 := Call(__e, PrimFunc(symshen_4_5_1out), W716)


__e.TailApply(tmp2330, tmp2333)
return


}


}, 1)

tmp2336 := Call(__e, PrimFunc(symshen_4_5pattern_6), V714)


tmp2337 := Call(__e, tmp2329, tmp2336)


__e.TailApply(tmp2326, tmp2337)
return


}, 1)

tmp2338 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp2325)


_ = tmp2338

tmp2339 := MakeNative(func(__e *ControlFlow) {
V719 := __e.Get(1)
_ = V719
tmp2340 := MakeNative(func(__e *ControlFlow) {
W720 := __e.Get(1)
_ = W720
tmp2341 := MakeNative(func(__e *ControlFlow) {
W721 := __e.Get(1)
_ = W721
tmp2342 := MakeNative(func(__e *ControlFlow) {
W722 := __e.Get(1)
_ = W722
__e.Return(W722)
return
}, 1)

tmp2343 := PrimStr(V719)

tmp2344 := Call(__e, PrimFunc(sym_8s), tmp2343, MakeString(")"))


tmp2345 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2344)


tmp2346 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2345)


tmp2347 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2346)


tmp2348 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp2347)


tmp2349 := PrimVectorSet(W721, MakeNumber(1), tmp2348)

__e.TailApply(tmp2342, tmp2349)
return


}, 1)

tmp2350 := PrimVectorSet(W720, MakeNumber(0), symshen_4printF)

__e.TailApply(tmp2341, tmp2350)
return


}, 1)

tmp2351 := PrimAbsvector(MakeNumber(2))

__e.TailApply(tmp2340, tmp2351)
return


}, 1)

tmp2352 := Call(__e, ns2_1set, symshen_4fn_1print, tmp2339)


_ = tmp2352

tmp2353 := MakeNative(func(__e *ControlFlow) {
V723 := __e.Get(1)
_ = V723
__e.Return(PrimVectorGet(V723, MakeNumber(1)))
return
}, 1)

tmp2354 := Call(__e, ns2_1set, symshen_4printF, tmp2353)


_ = tmp2354

tmp2355 := MakeNative(func(__e *ControlFlow) {
V728 := __e.Get(1)
_ = V728
V729 := __e.Get(2)
_ = V729
tmp2379 := PrimIsPair(V729)

var ifres2375 Obj

if True == tmp2379 {
tmp2377 := PrimTail(V729)

tmp2378 := PrimEqual(Nil, tmp2377)

var ifres2376 Obj

if True == tmp2378 {
ifres2376 = True


} else {
ifres2376 = False


}

ifres2375 = ifres2376


} else {
ifres2375 = False


}

if True == ifres2375 {
tmp2356 := PrimHead(V729)

__e.TailApply(PrimFunc(symlength), tmp2356)
return


} else {
tmp2373 := PrimIsPair(V729)

var ifres2361 Obj

if True == tmp2373 {
tmp2371 := PrimTail(V729)

tmp2372 := PrimIsPair(tmp2371)

var ifres2363 Obj

if True == tmp2372 {
tmp2365 := PrimHead(V729)

tmp2366 := Call(__e, PrimFunc(symlength), tmp2365)


tmp2367 := PrimTail(V729)

tmp2368 := PrimHead(tmp2367)

tmp2369 := Call(__e, PrimFunc(symlength), tmp2368)


tmp2370 := PrimEqual(tmp2366, tmp2369)

var ifres2364 Obj

if True == tmp2370 {
ifres2364 = True


} else {
ifres2364 = False


}

ifres2363 = ifres2364


} else {
ifres2363 = False


}

var ifres2362 Obj

if True == ifres2363 {
ifres2362 = True


} else {
ifres2362 = False


}

ifres2361 = ifres2362


} else {
ifres2361 = False


}

if True == ifres2361 {
tmp2357 := PrimTail(V729)

__e.TailApply(PrimFunc(symshen_4arity_1chk), V728, tmp2357)
return


} else {
tmp2358 := Call(__e, PrimFunc(symshen_4app), V728, MakeString("\n"), symshen_4a)


tmp2359 := PrimStringConcat(MakeString("arity error in "), tmp2358)

__e.Return(PrimSimpleError(tmp2359))
return


}


}


}, 2)

tmp2380 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp2355)


_ = tmp2380

tmp2381 := MakeNative(func(__e *ControlFlow) {
V730 := __e.Get(1)
_ = V730
V731 := __e.Get(2)
_ = V731
tmp2387 := Call(__e, PrimFunc(symtuple_2), V731)


if True == tmp2387 {
tmp2382 := Call(__e, PrimFunc(symfst), V731)


tmp2383 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp2382)


tmp2384 := Call(__e, PrimFunc(symsnd), V731)


tmp2385 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp2383, tmp2384)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V730, tmp2385)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4free_1var_1chk)
return
}


}, 2)

tmp2388 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp2381)


_ = tmp2388

tmp2389 := MakeNative(func(__e *ControlFlow) {
V732 := __e.Get(1)
_ = V732
V733 := __e.Get(2)
_ = V733
tmp2401 := Call(__e, PrimFunc(symempty_2), V733)


if True == tmp2401 {
__e.Return(symshen_4skip)
return
} else {
tmp2390 := Call(__e, PrimFunc(symshen_4app), V732, MakeString(":"), symshen_4a)


tmp2391 := PrimStringConcat(MakeString("free variables in "), tmp2390)

tmp2392 := Call(__e, PrimFunc(symstoutput))


tmp2393 := Call(__e, PrimFunc(sympr), tmp2391, tmp2392)


_ = tmp2393

tmp2394 := MakeNative(func(__e *ControlFlow) {
Z734 := __e.Get(1)
_ = Z734
tmp2395 := Call(__e, PrimFunc(symshen_4app), Z734, MakeString(""), symshen_4a)


tmp2396 := PrimStringConcat(MakeString(" "), tmp2395)

tmp2397 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp2396, tmp2397)
return


}, 1)

tmp2398 := Call(__e, PrimFunc(symshen_4for_1each), tmp2394, V733)


_ = tmp2398

tmp2399 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp2399

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp2402 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp2389)


_ = tmp2402

tmp2403 := MakeNative(func(__e *ControlFlow) {
V737 := __e.Get(1)
_ = V737
tmp2411 := PrimIsVariable(V737)

if True == tmp2411 {
__e.Return(PrimCons(V737, Nil))
return
} else {
tmp2409 := PrimIsPair(V737)

if True == tmp2409 {
tmp2404 := PrimHead(V737)

tmp2405 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp2404)


tmp2406 := PrimTail(V737)

tmp2407 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp2406)


__e.TailApply(PrimFunc(symunion), tmp2405, tmp2407)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp2412 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp2403)


_ = tmp2412

tmp2413 := MakeNative(func(__e *ControlFlow) {
V742 := __e.Get(1)
_ = V742
V743 := __e.Get(2)
_ = V743
tmp2503 := PrimIsPair(V743)

var ifres2490 Obj

if True == tmp2503 {
tmp2501 := PrimHead(V743)

tmp2502 := PrimEqual(symprotect, tmp2501)

var ifres2492 Obj

if True == tmp2502 {
tmp2499 := PrimTail(V743)

tmp2500 := PrimIsPair(tmp2499)

var ifres2494 Obj

if True == tmp2500 {
tmp2496 := PrimTail(V743)

tmp2497 := PrimTail(tmp2496)

tmp2498 := PrimEqual(Nil, tmp2497)

var ifres2495 Obj

if True == tmp2498 {
ifres2495 = True


} else {
ifres2495 = False


}

ifres2494 = ifres2495


} else {
ifres2494 = False


}

var ifres2493 Obj

if True == ifres2494 {
ifres2493 = True


} else {
ifres2493 = False


}

ifres2492 = ifres2493


} else {
ifres2492 = False


}

var ifres2491 Obj

if True == ifres2492 {
ifres2491 = True


} else {
ifres2491 = False


}

ifres2490 = ifres2491


} else {
ifres2490 = False


}

if True == ifres2490 {
__e.Return(Nil)
return
} else {
tmp2488 := PrimIsPair(V743)

var ifres2462 Obj

if True == tmp2488 {
tmp2486 := PrimHead(V743)

tmp2487 := PrimEqual(symlet, tmp2486)

var ifres2464 Obj

if True == tmp2487 {
tmp2484 := PrimTail(V743)

tmp2485 := PrimIsPair(tmp2484)

var ifres2466 Obj

if True == tmp2485 {
tmp2481 := PrimTail(V743)

tmp2482 := PrimTail(tmp2481)

tmp2483 := PrimIsPair(tmp2482)

var ifres2468 Obj

if True == tmp2483 {
tmp2477 := PrimTail(V743)

tmp2478 := PrimTail(tmp2477)

tmp2479 := PrimTail(tmp2478)

tmp2480 := PrimIsPair(tmp2479)

var ifres2470 Obj

if True == tmp2480 {
tmp2472 := PrimTail(V743)

tmp2473 := PrimTail(tmp2472)

tmp2474 := PrimTail(tmp2473)

tmp2475 := PrimTail(tmp2474)

tmp2476 := PrimEqual(Nil, tmp2475)

var ifres2471 Obj

if True == tmp2476 {
ifres2471 = True


} else {
ifres2471 = False


}

ifres2470 = ifres2471


} else {
ifres2470 = False


}

var ifres2469 Obj

if True == ifres2470 {
ifres2469 = True


} else {
ifres2469 = False


}

ifres2468 = ifres2469


} else {
ifres2468 = False


}

var ifres2467 Obj

if True == ifres2468 {
ifres2467 = True


} else {
ifres2467 = False


}

ifres2466 = ifres2467


} else {
ifres2466 = False


}

var ifres2465 Obj

if True == ifres2466 {
ifres2465 = True


} else {
ifres2465 = False


}

ifres2464 = ifres2465


} else {
ifres2464 = False


}

var ifres2463 Obj

if True == ifres2464 {
ifres2463 = True


} else {
ifres2463 = False


}

ifres2462 = ifres2463


} else {
ifres2462 = False


}

if True == ifres2462 {
tmp2414 := PrimTail(V743)

tmp2415 := PrimTail(tmp2414)

tmp2416 := PrimHead(tmp2415)

tmp2417 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V742, tmp2416)


tmp2418 := PrimTail(V743)

tmp2419 := PrimHead(tmp2418)

tmp2420 := PrimCons(tmp2419, V742)

tmp2421 := PrimTail(V743)

tmp2422 := PrimTail(tmp2421)

tmp2423 := PrimTail(tmp2422)

tmp2424 := PrimHead(tmp2423)

tmp2425 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp2420, tmp2424)


__e.TailApply(PrimFunc(symunion), tmp2417, tmp2425)
return


} else {
tmp2460 := PrimIsPair(V743)

var ifres2441 Obj

if True == tmp2460 {
tmp2458 := PrimHead(V743)

tmp2459 := PrimEqual(symlambda, tmp2458)

var ifres2443 Obj

if True == tmp2459 {
tmp2456 := PrimTail(V743)

tmp2457 := PrimIsPair(tmp2456)

var ifres2445 Obj

if True == tmp2457 {
tmp2453 := PrimTail(V743)

tmp2454 := PrimTail(tmp2453)

tmp2455 := PrimIsPair(tmp2454)

var ifres2447 Obj

if True == tmp2455 {
tmp2449 := PrimTail(V743)

tmp2450 := PrimTail(tmp2449)

tmp2451 := PrimTail(tmp2450)

tmp2452 := PrimEqual(Nil, tmp2451)

var ifres2448 Obj

if True == tmp2452 {
ifres2448 = True


} else {
ifres2448 = False


}

ifres2447 = ifres2448


} else {
ifres2447 = False


}

var ifres2446 Obj

if True == ifres2447 {
ifres2446 = True


} else {
ifres2446 = False


}

ifres2445 = ifres2446


} else {
ifres2445 = False


}

var ifres2444 Obj

if True == ifres2445 {
ifres2444 = True


} else {
ifres2444 = False


}

ifres2443 = ifres2444


} else {
ifres2443 = False


}

var ifres2442 Obj

if True == ifres2443 {
ifres2442 = True


} else {
ifres2442 = False


}

ifres2441 = ifres2442


} else {
ifres2441 = False


}

if True == ifres2441 {
tmp2426 := PrimTail(V743)

tmp2427 := PrimHead(tmp2426)

tmp2428 := PrimCons(tmp2427, V742)

tmp2429 := PrimTail(V743)

tmp2430 := PrimTail(tmp2429)

tmp2431 := PrimHead(tmp2430)

__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp2428, tmp2431)
return


} else {
tmp2439 := PrimIsPair(V743)

if True == tmp2439 {
tmp2432 := PrimHead(V743)

tmp2433 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V742, tmp2432)


tmp2434 := PrimTail(V743)

tmp2435 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V742, tmp2434)


__e.TailApply(PrimFunc(symunion), tmp2433, tmp2435)
return


} else {
tmp2437 := Call(__e, PrimFunc(symshen_4free_1variable_2), V743, V742)


if True == tmp2437 {
__e.Return(PrimCons(V743, Nil))
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

tmp2504 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp2413)


_ = tmp2504

tmp2505 := MakeNative(func(__e *ControlFlow) {
V744 := __e.Get(1)
_ = V744
V745 := __e.Get(2)
_ = V745
tmp2510 := PrimIsVariable(V744)

if True == tmp2510 {
tmp2507 := Call(__e, PrimFunc(symelement_2), V744, V745)


tmp2508 := PrimNot(tmp2507)

if True == tmp2508 {
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

tmp2511 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp2505)


_ = tmp2511

tmp2512 := MakeNative(func(__e *ControlFlow) {
V746 := __e.Get(1)
_ = V746
V747 := __e.Get(2)
_ = V747
tmp2513 := PrimValue(symshen_4_duserdefs_d)

tmp2514 := Call(__e, PrimFunc(symadjoin), V746, tmp2513)


tmp2515 := PrimSet(symshen_4_duserdefs_d, tmp2514)

_ = tmp2515

tmp2516 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V746, symshen_4source, V747, tmp2516)
return


}, 2)

tmp2517 := Call(__e, ns2_1set, symshen_4record_1kl, tmp2512)


_ = tmp2517

tmp2518 := MakeNative(func(__e *ControlFlow) {
V748 := __e.Get(1)
_ = V748
V749 := __e.Get(2)
_ = V749
V750 := __e.Get(3)
_ = V750
tmp2519 := MakeNative(func(__e *ControlFlow) {
W751 := __e.Get(1)
_ = W751
tmp2520 := MakeNative(func(__e *ControlFlow) {
W752 := __e.Get(1)
_ = W752
tmp2521 := MakeNative(func(__e *ControlFlow) {
W753 := __e.Get(1)
_ = W753
__e.Return(W753)
return
}, 1)

tmp2522 := Call(__e, PrimFunc(symshen_4cond_1form), W752)


tmp2523 := PrimCons(tmp2522, Nil)

tmp2524 := PrimCons(W751, tmp2523)

tmp2525 := PrimCons(V748, tmp2524)

tmp2526 := PrimCons(symdefun, tmp2525)

__e.TailApply(tmp2521, tmp2526)
return


}, 1)

tmp2527 := Call(__e, PrimFunc(symshen_4kl_1body), V749, W751)


tmp2528 := Call(__e, PrimFunc(symshen_4scan_1body), V748, tmp2527)


__e.TailApply(tmp2520, tmp2528)
return


}, 1)

tmp2529 := Call(__e, PrimFunc(symshen_4parameters), V750)


__e.TailApply(tmp2519, tmp2529)
return


}, 3)

tmp2530 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp2518)


_ = tmp2530

tmp2531 := MakeNative(func(__e *ControlFlow) {
V754 := __e.Get(1)
_ = V754
tmp2536 := PrimEqual(MakeNumber(0), V754)

if True == tmp2536 {
__e.Return(Nil)
return
} else {
tmp2532 := Call(__e, PrimFunc(symgensym), symV)


tmp2533 := PrimNumberSubtract(V754, MakeNumber(1))

tmp2534 := Call(__e, PrimFunc(symshen_4parameters), tmp2533)


__e.Return(PrimCons(tmp2532, tmp2534))
return


}


}, 1)

tmp2537 := Call(__e, ns2_1set, symshen_4parameters, tmp2531)


_ = tmp2537

tmp2538 := MakeNative(func(__e *ControlFlow) {
V757 := __e.Get(1)
_ = V757
tmp2562 := PrimIsPair(V757)

var ifres2542 Obj

if True == tmp2562 {
tmp2560 := PrimHead(V757)

tmp2561 := PrimIsPair(tmp2560)

var ifres2544 Obj

if True == tmp2561 {
tmp2557 := PrimHead(V757)

tmp2558 := PrimHead(tmp2557)

tmp2559 := PrimEqual(True, tmp2558)

var ifres2546 Obj

if True == tmp2559 {
tmp2554 := PrimHead(V757)

tmp2555 := PrimTail(tmp2554)

tmp2556 := PrimIsPair(tmp2555)

var ifres2548 Obj

if True == tmp2556 {
tmp2550 := PrimHead(V757)

tmp2551 := PrimTail(tmp2550)

tmp2552 := PrimTail(tmp2551)

tmp2553 := PrimEqual(Nil, tmp2552)

var ifres2549 Obj

if True == tmp2553 {
ifres2549 = True


} else {
ifres2549 = False


}

ifres2548 = ifres2549


} else {
ifres2548 = False


}

var ifres2547 Obj

if True == ifres2548 {
ifres2547 = True


} else {
ifres2547 = False


}

ifres2546 = ifres2547


} else {
ifres2546 = False


}

var ifres2545 Obj

if True == ifres2546 {
ifres2545 = True


} else {
ifres2545 = False


}

ifres2544 = ifres2545


} else {
ifres2544 = False


}

var ifres2543 Obj

if True == ifres2544 {
ifres2543 = True


} else {
ifres2543 = False


}

ifres2542 = ifres2543


} else {
ifres2542 = False


}

if True == ifres2542 {
tmp2539 := PrimHead(V757)

tmp2540 := PrimTail(tmp2539)

__e.Return(PrimHead(tmp2540))
return


} else {
__e.Return(PrimCons(symcond, V757))
return
}


}, 1)

tmp2563 := Call(__e, ns2_1set, symshen_4cond_1form, tmp2538)


_ = tmp2563

tmp2564 := MakeNative(func(__e *ControlFlow) {
V766 := __e.Get(1)
_ = V766
V767 := __e.Get(2)
_ = V767
tmp2608 := PrimEqual(Nil, V767)

if True == tmp2608 {
tmp2565 := PrimCons(V766, Nil)

tmp2566 := PrimCons(symshen_4f_1error, tmp2565)

tmp2567 := PrimCons(tmp2566, Nil)

tmp2568 := PrimCons(True, tmp2567)

__e.Return(PrimCons(tmp2568, Nil))
return


} else {
tmp2606 := PrimIsPair(V767)

var ifres2602 Obj

if True == tmp2606 {
tmp2604 := PrimHead(V767)

tmp2605 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp2604)


var ifres2603 Obj

if True == tmp2605 {
ifres2603 = True


} else {
ifres2603 = False


}

ifres2602 = ifres2603


} else {
ifres2602 = False


}

if True == ifres2602 {
tmp2569 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp2570 := Call(__e, PrimFunc(symgensym), symResult)


tmp2571 := PrimHead(V767)

tmp2572 := PrimTail(V767)

__e.TailApply(PrimFunc(symshen_4choicepoint), V766, tmp2569, tmp2570, tmp2571, tmp2572)
return


} else {
tmp2600 := PrimIsPair(V767)

var ifres2580 Obj

if True == tmp2600 {
tmp2598 := PrimHead(V767)

tmp2599 := PrimIsPair(tmp2598)

var ifres2582 Obj

if True == tmp2599 {
tmp2595 := PrimHead(V767)

tmp2596 := PrimHead(tmp2595)

tmp2597 := PrimEqual(True, tmp2596)

var ifres2584 Obj

if True == tmp2597 {
tmp2592 := PrimHead(V767)

tmp2593 := PrimTail(tmp2592)

tmp2594 := PrimIsPair(tmp2593)

var ifres2586 Obj

if True == tmp2594 {
tmp2588 := PrimHead(V767)

tmp2589 := PrimTail(tmp2588)

tmp2590 := PrimTail(tmp2589)

tmp2591 := PrimEqual(Nil, tmp2590)

var ifres2587 Obj

if True == tmp2591 {
ifres2587 = True


} else {
ifres2587 = False


}

ifres2586 = ifres2587


} else {
ifres2586 = False


}

var ifres2585 Obj

if True == ifres2586 {
ifres2585 = True


} else {
ifres2585 = False


}

ifres2584 = ifres2585


} else {
ifres2584 = False


}

var ifres2583 Obj

if True == ifres2584 {
ifres2583 = True


} else {
ifres2583 = False


}

ifres2582 = ifres2583


} else {
ifres2582 = False


}

var ifres2581 Obj

if True == ifres2582 {
ifres2581 = True


} else {
ifres2581 = False


}

ifres2580 = ifres2581


} else {
ifres2580 = False


}

if True == ifres2580 {
tmp2573 := PrimHead(V767)

__e.Return(PrimCons(tmp2573, Nil))
return


} else {
tmp2578 := PrimIsPair(V767)

if True == tmp2578 {
tmp2574 := PrimHead(V767)

tmp2575 := PrimTail(V767)

tmp2576 := Call(__e, PrimFunc(symshen_4scan_1body), V766, tmp2575)


__e.Return(PrimCons(tmp2574, tmp2576))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.scan-body")))
return
}


}


}


}


}, 2)

tmp2609 := Call(__e, ns2_1set, symshen_4scan_1body, tmp2564)


_ = tmp2609

tmp2610 := MakeNative(func(__e *ControlFlow) {
V774 := __e.Get(1)
_ = V774
tmp2645 := PrimIsPair(V774)

var ifres2612 Obj

if True == tmp2645 {
tmp2643 := PrimTail(V774)

tmp2644 := PrimIsPair(tmp2643)

var ifres2614 Obj

if True == tmp2644 {
tmp2640 := PrimTail(V774)

tmp2641 := PrimHead(tmp2640)

tmp2642 := PrimIsPair(tmp2641)

var ifres2616 Obj

if True == tmp2642 {
tmp2636 := PrimTail(V774)

tmp2637 := PrimHead(tmp2636)

tmp2638 := PrimHead(tmp2637)

tmp2639 := PrimEqual(symshen_4choicepoint_b, tmp2638)

var ifres2618 Obj

if True == tmp2639 {
tmp2632 := PrimTail(V774)

tmp2633 := PrimHead(tmp2632)

tmp2634 := PrimTail(tmp2633)

tmp2635 := PrimIsPair(tmp2634)

var ifres2620 Obj

if True == tmp2635 {
tmp2627 := PrimTail(V774)

tmp2628 := PrimHead(tmp2627)

tmp2629 := PrimTail(tmp2628)

tmp2630 := PrimTail(tmp2629)

tmp2631 := PrimEqual(Nil, tmp2630)

var ifres2622 Obj

if True == tmp2631 {
tmp2624 := PrimTail(V774)

tmp2625 := PrimTail(tmp2624)

tmp2626 := PrimEqual(Nil, tmp2625)

var ifres2623 Obj

if True == tmp2626 {
ifres2623 = True


} else {
ifres2623 = False


}

ifres2622 = ifres2623


} else {
ifres2622 = False


}

var ifres2621 Obj

if True == ifres2622 {
ifres2621 = True


} else {
ifres2621 = False


}

ifres2620 = ifres2621


} else {
ifres2620 = False


}

var ifres2619 Obj

if True == ifres2620 {
ifres2619 = True


} else {
ifres2619 = False


}

ifres2618 = ifres2619


} else {
ifres2618 = False


}

var ifres2617 Obj

if True == ifres2618 {
ifres2617 = True


} else {
ifres2617 = False


}

ifres2616 = ifres2617


} else {
ifres2616 = False


}

var ifres2615 Obj

if True == ifres2616 {
ifres2615 = True


} else {
ifres2615 = False


}

ifres2614 = ifres2615


} else {
ifres2614 = False


}

var ifres2613 Obj

if True == ifres2614 {
ifres2613 = True


} else {
ifres2613 = False


}

ifres2612 = ifres2613


} else {
ifres2612 = False


}

if True == ifres2612 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp2646 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp2610)


_ = tmp2646

tmp2647 := MakeNative(func(__e *ControlFlow) {
V790 := __e.Get(1)
_ = V790
V791 := __e.Get(2)
_ = V791
V792 := __e.Get(3)
_ = V792
V793 := __e.Get(4)
_ = V793
V794 := __e.Get(5)
_ = V794
tmp2839 := PrimIsPair(V793)

var ifres2761 Obj

if True == tmp2839 {
tmp2837 := PrimTail(V793)

tmp2838 := PrimIsPair(tmp2837)

var ifres2763 Obj

if True == tmp2838 {
tmp2834 := PrimTail(V793)

tmp2835 := PrimHead(tmp2834)

tmp2836 := PrimIsPair(tmp2835)

var ifres2765 Obj

if True == tmp2836 {
tmp2830 := PrimTail(V793)

tmp2831 := PrimHead(tmp2830)

tmp2832 := PrimTail(tmp2831)

tmp2833 := PrimIsPair(tmp2832)

var ifres2767 Obj

if True == tmp2833 {
tmp2825 := PrimTail(V793)

tmp2826 := PrimHead(tmp2825)

tmp2827 := PrimTail(tmp2826)

tmp2828 := PrimHead(tmp2827)

tmp2829 := PrimIsPair(tmp2828)

var ifres2769 Obj

if True == tmp2829 {
tmp2819 := PrimTail(V793)

tmp2820 := PrimHead(tmp2819)

tmp2821 := PrimTail(tmp2820)

tmp2822 := PrimHead(tmp2821)

tmp2823 := PrimHead(tmp2822)

tmp2824 := PrimEqual(symfail_1if, tmp2823)

var ifres2771 Obj

if True == tmp2824 {
tmp2813 := PrimTail(V793)

tmp2814 := PrimHead(tmp2813)

tmp2815 := PrimTail(tmp2814)

tmp2816 := PrimHead(tmp2815)

tmp2817 := PrimTail(tmp2816)

tmp2818 := PrimIsPair(tmp2817)

var ifres2773 Obj

if True == tmp2818 {
tmp2806 := PrimTail(V793)

tmp2807 := PrimHead(tmp2806)

tmp2808 := PrimTail(tmp2807)

tmp2809 := PrimHead(tmp2808)

tmp2810 := PrimTail(tmp2809)

tmp2811 := PrimTail(tmp2810)

tmp2812 := PrimIsPair(tmp2811)

var ifres2775 Obj

if True == tmp2812 {
tmp2798 := PrimTail(V793)

tmp2799 := PrimHead(tmp2798)

tmp2800 := PrimTail(tmp2799)

tmp2801 := PrimHead(tmp2800)

tmp2802 := PrimTail(tmp2801)

tmp2803 := PrimTail(tmp2802)

tmp2804 := PrimTail(tmp2803)

tmp2805 := PrimEqual(Nil, tmp2804)

var ifres2777 Obj

if True == tmp2805 {
tmp2793 := PrimTail(V793)

tmp2794 := PrimHead(tmp2793)

tmp2795 := PrimTail(tmp2794)

tmp2796 := PrimTail(tmp2795)

tmp2797 := PrimEqual(Nil, tmp2796)

var ifres2779 Obj

if True == tmp2797 {
tmp2790 := PrimTail(V793)

tmp2791 := PrimTail(tmp2790)

tmp2792 := PrimEqual(Nil, tmp2791)

var ifres2781 Obj

if True == tmp2792 {
tmp2783 := PrimTail(V793)

tmp2784 := PrimHead(tmp2783)

tmp2785 := PrimTail(tmp2784)

tmp2786 := PrimHead(tmp2785)

tmp2787 := PrimTail(tmp2786)

tmp2788 := PrimHead(tmp2787)

tmp2789 := PrimEqual(V790, tmp2788)

var ifres2782 Obj

if True == tmp2789 {
ifres2782 = True


} else {
ifres2782 = False


}

ifres2781 = ifres2782


} else {
ifres2781 = False


}

var ifres2780 Obj

if True == ifres2781 {
ifres2780 = True


} else {
ifres2780 = False


}

ifres2779 = ifres2780


} else {
ifres2779 = False


}

var ifres2778 Obj

if True == ifres2779 {
ifres2778 = True


} else {
ifres2778 = False


}

ifres2777 = ifres2778


} else {
ifres2777 = False


}

var ifres2776 Obj

if True == ifres2777 {
ifres2776 = True


} else {
ifres2776 = False


}

ifres2775 = ifres2776


} else {
ifres2775 = False


}

var ifres2774 Obj

if True == ifres2775 {
ifres2774 = True


} else {
ifres2774 = False


}

ifres2773 = ifres2774


} else {
ifres2773 = False


}

var ifres2772 Obj

if True == ifres2773 {
ifres2772 = True


} else {
ifres2772 = False


}

ifres2771 = ifres2772


} else {
ifres2771 = False


}

var ifres2770 Obj

if True == ifres2771 {
ifres2770 = True


} else {
ifres2770 = False


}

ifres2769 = ifres2770


} else {
ifres2769 = False


}

var ifres2768 Obj

if True == ifres2769 {
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

if True == ifres2761 {
tmp2648 := PrimTail(V793)

tmp2649 := PrimHead(tmp2648)

tmp2650 := PrimTail(tmp2649)

tmp2651 := PrimHead(tmp2650)

tmp2652 := PrimTail(tmp2651)

tmp2653 := PrimHead(tmp2652)

tmp2654 := Call(__e, PrimFunc(symshen_4scan_1body), tmp2653, V794)


tmp2655 := PrimCons(symcond, tmp2654)

tmp2656 := PrimCons(tmp2655, Nil)

tmp2657 := PrimCons(symfreeze, tmp2656)

tmp2658 := PrimHead(V793)

tmp2659 := PrimTail(V793)

tmp2660 := PrimHead(tmp2659)

tmp2661 := PrimTail(tmp2660)

tmp2662 := PrimHead(tmp2661)

tmp2663 := PrimTail(tmp2662)

tmp2664 := PrimTail(tmp2663)

tmp2665 := PrimHead(tmp2664)

tmp2666 := PrimTail(V793)

tmp2667 := PrimHead(tmp2666)

tmp2668 := PrimTail(tmp2667)

tmp2669 := PrimHead(tmp2668)

tmp2670 := PrimTail(tmp2669)

tmp2671 := PrimHead(tmp2670)

tmp2672 := PrimCons(V792, Nil)

tmp2673 := PrimCons(tmp2671, tmp2672)

tmp2674 := PrimCons(V791, Nil)

tmp2675 := PrimCons(symthaw, tmp2674)

tmp2676 := PrimCons(V792, Nil)

tmp2677 := PrimCons(tmp2675, tmp2676)

tmp2678 := PrimCons(tmp2673, tmp2677)

tmp2679 := PrimCons(symif, tmp2678)

tmp2680 := PrimCons(tmp2679, Nil)

tmp2681 := PrimCons(tmp2665, tmp2680)

tmp2682 := PrimCons(V792, tmp2681)

tmp2683 := PrimCons(symlet, tmp2682)

tmp2684 := PrimCons(V791, Nil)

tmp2685 := PrimCons(symthaw, tmp2684)

tmp2686 := PrimCons(tmp2685, Nil)

tmp2687 := PrimCons(tmp2683, tmp2686)

tmp2688 := PrimCons(tmp2658, tmp2687)

tmp2689 := PrimCons(symif, tmp2688)

tmp2690 := PrimCons(tmp2689, Nil)

tmp2691 := PrimCons(tmp2657, tmp2690)

tmp2692 := PrimCons(V791, tmp2691)

tmp2693 := PrimCons(symlet, tmp2692)

tmp2694 := PrimCons(tmp2693, Nil)

tmp2695 := PrimCons(True, tmp2694)

__e.Return(PrimCons(tmp2695, Nil))
return


} else {
tmp2759 := PrimIsPair(V793)

var ifres2732 Obj

if True == tmp2759 {
tmp2757 := PrimTail(V793)

tmp2758 := PrimIsPair(tmp2757)

var ifres2734 Obj

if True == tmp2758 {
tmp2754 := PrimTail(V793)

tmp2755 := PrimHead(tmp2754)

tmp2756 := PrimIsPair(tmp2755)

var ifres2736 Obj

if True == tmp2756 {
tmp2750 := PrimTail(V793)

tmp2751 := PrimHead(tmp2750)

tmp2752 := PrimTail(tmp2751)

tmp2753 := PrimIsPair(tmp2752)

var ifres2738 Obj

if True == tmp2753 {
tmp2745 := PrimTail(V793)

tmp2746 := PrimHead(tmp2745)

tmp2747 := PrimTail(tmp2746)

tmp2748 := PrimTail(tmp2747)

tmp2749 := PrimEqual(Nil, tmp2748)

var ifres2740 Obj

if True == tmp2749 {
tmp2742 := PrimTail(V793)

tmp2743 := PrimTail(tmp2742)

tmp2744 := PrimEqual(Nil, tmp2743)

var ifres2741 Obj

if True == tmp2744 {
ifres2741 = True


} else {
ifres2741 = False


}

ifres2740 = ifres2741


} else {
ifres2740 = False


}

var ifres2739 Obj

if True == ifres2740 {
ifres2739 = True


} else {
ifres2739 = False


}

ifres2738 = ifres2739


} else {
ifres2738 = False


}

var ifres2737 Obj

if True == ifres2738 {
ifres2737 = True


} else {
ifres2737 = False


}

ifres2736 = ifres2737


} else {
ifres2736 = False


}

var ifres2735 Obj

if True == ifres2736 {
ifres2735 = True


} else {
ifres2735 = False


}

ifres2734 = ifres2735


} else {
ifres2734 = False


}

var ifres2733 Obj

if True == ifres2734 {
ifres2733 = True


} else {
ifres2733 = False


}

ifres2732 = ifres2733


} else {
ifres2732 = False


}

if True == ifres2732 {
tmp2696 := Call(__e, PrimFunc(symshen_4scan_1body), V790, V794)


tmp2697 := PrimCons(symcond, tmp2696)

tmp2698 := PrimCons(tmp2697, Nil)

tmp2699 := PrimCons(symfreeze, tmp2698)

tmp2700 := PrimHead(V793)

tmp2701 := PrimTail(V793)

tmp2702 := PrimHead(tmp2701)

tmp2703 := PrimTail(tmp2702)

tmp2704 := PrimHead(tmp2703)

tmp2705 := PrimCons(symfail, Nil)

tmp2706 := PrimCons(tmp2705, Nil)

tmp2707 := PrimCons(V792, tmp2706)

tmp2708 := PrimCons(sym_a, tmp2707)

tmp2709 := PrimCons(V791, Nil)

tmp2710 := PrimCons(symthaw, tmp2709)

tmp2711 := PrimCons(V792, Nil)

tmp2712 := PrimCons(tmp2710, tmp2711)

tmp2713 := PrimCons(tmp2708, tmp2712)

tmp2714 := PrimCons(symif, tmp2713)

tmp2715 := PrimCons(tmp2714, Nil)

tmp2716 := PrimCons(tmp2704, tmp2715)

tmp2717 := PrimCons(V792, tmp2716)

tmp2718 := PrimCons(symlet, tmp2717)

tmp2719 := PrimCons(V791, Nil)

tmp2720 := PrimCons(symthaw, tmp2719)

tmp2721 := PrimCons(tmp2720, Nil)

tmp2722 := PrimCons(tmp2718, tmp2721)

tmp2723 := PrimCons(tmp2700, tmp2722)

tmp2724 := PrimCons(symif, tmp2723)

tmp2725 := PrimCons(tmp2724, Nil)

tmp2726 := PrimCons(tmp2699, tmp2725)

tmp2727 := PrimCons(V791, tmp2726)

tmp2728 := PrimCons(symlet, tmp2727)

tmp2729 := PrimCons(tmp2728, Nil)

tmp2730 := PrimCons(True, tmp2729)

__e.Return(PrimCons(tmp2730, Nil))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.choicepoint")))
return
}


}


}, 5)

tmp2840 := Call(__e, ns2_1set, symshen_4choicepoint, tmp2647)


_ = tmp2840

tmp2841 := MakeNative(func(__e *ControlFlow) {
V796 := __e.Get(1)
_ = V796
V797 := __e.Get(2)
_ = V797
V798 := __e.Get(3)
_ = V798
tmp2855 := PrimEqual(V796, V798)

if True == tmp2855 {
__e.Return(V797)
return
} else {
tmp2853 := PrimIsPair(V798)

if True == tmp2853 {
tmp2842 := MakeNative(func(__e *ControlFlow) {
W799 := __e.Get(1)
_ = W799
tmp2848 := PrimHead(V798)

tmp2849 := PrimEqual(W799, tmp2848)

if True == tmp2849 {
tmp2843 := PrimHead(V798)

tmp2844 := PrimTail(V798)

tmp2845 := Call(__e, PrimFunc(symshen_4rep_1X), V796, V797, tmp2844)


__e.Return(PrimCons(tmp2843, tmp2845))
return


} else {
tmp2846 := PrimTail(V798)

__e.Return(PrimCons(W799, tmp2846))
return


}


}, 1)

tmp2850 := PrimHead(V798)

tmp2851 := Call(__e, PrimFunc(symshen_4rep_1X), V796, V797, tmp2850)


__e.TailApply(tmp2842, tmp2851)
return


} else {
__e.Return(V798)
return
}


}


}, 3)

tmp2856 := Call(__e, ns2_1set, symshen_4rep_1X, tmp2841)


_ = tmp2856

tmp2857 := MakeNative(func(__e *ControlFlow) {
V800 := __e.Get(1)
_ = V800
tmp2940 := PrimIsPair(V800)

var ifres2921 Obj

if True == tmp2940 {
tmp2938 := PrimHead(V800)

tmp2939 := PrimEqual(symlambda, tmp2938)

var ifres2923 Obj

if True == tmp2939 {
tmp2936 := PrimTail(V800)

tmp2937 := PrimIsPair(tmp2936)

var ifres2925 Obj

if True == tmp2937 {
tmp2933 := PrimTail(V800)

tmp2934 := PrimTail(tmp2933)

tmp2935 := PrimIsPair(tmp2934)

var ifres2927 Obj

if True == tmp2935 {
tmp2929 := PrimTail(V800)

tmp2930 := PrimTail(tmp2929)

tmp2931 := PrimTail(tmp2930)

tmp2932 := PrimEqual(Nil, tmp2931)

var ifres2928 Obj

if True == tmp2932 {
ifres2928 = True


} else {
ifres2928 = False


}

ifres2927 = ifres2928


} else {
ifres2927 = False


}

var ifres2926 Obj

if True == ifres2927 {
ifres2926 = True


} else {
ifres2926 = False


}

ifres2925 = ifres2926


} else {
ifres2925 = False


}

var ifres2924 Obj

if True == ifres2925 {
ifres2924 = True


} else {
ifres2924 = False


}

ifres2923 = ifres2924


} else {
ifres2923 = False


}

var ifres2922 Obj

if True == ifres2923 {
ifres2922 = True


} else {
ifres2922 = False


}

ifres2921 = ifres2922


} else {
ifres2921 = False


}

if True == ifres2921 {
tmp2858 := MakeNative(func(__e *ControlFlow) {
W801 := __e.Get(1)
_ = W801
tmp2859 := MakeNative(func(__e *ControlFlow) {
W802 := __e.Get(1)
_ = W802
tmp2860 := MakeNative(func(__e *ControlFlow) {
Z803 := __e.Get(1)
_ = Z803
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z803)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2860, W802)
return


}, 1)

tmp2861 := PrimTail(V800)

tmp2862 := PrimHead(tmp2861)

tmp2863 := PrimTail(V800)

tmp2864 := PrimTail(tmp2863)

tmp2865 := PrimHead(tmp2864)

tmp2866 := Call(__e, PrimFunc(symshen_4beta), tmp2862, W801, tmp2865)


tmp2867 := PrimCons(tmp2866, Nil)

tmp2868 := PrimCons(W801, tmp2867)

tmp2869 := PrimCons(symlambda, tmp2868)

__e.TailApply(tmp2859, tmp2869)
return


}, 1)

tmp2870 := Call(__e, PrimFunc(symgensym), symZ)


__e.TailApply(tmp2858, tmp2870)
return


} else {
tmp2919 := PrimIsPair(V800)

var ifres2893 Obj

if True == tmp2919 {
tmp2917 := PrimHead(V800)

tmp2918 := PrimEqual(symlet, tmp2917)

var ifres2895 Obj

if True == tmp2918 {
tmp2915 := PrimTail(V800)

tmp2916 := PrimIsPair(tmp2915)

var ifres2897 Obj

if True == tmp2916 {
tmp2912 := PrimTail(V800)

tmp2913 := PrimTail(tmp2912)

tmp2914 := PrimIsPair(tmp2913)

var ifres2899 Obj

if True == tmp2914 {
tmp2908 := PrimTail(V800)

tmp2909 := PrimTail(tmp2908)

tmp2910 := PrimTail(tmp2909)

tmp2911 := PrimIsPair(tmp2910)

var ifres2901 Obj

if True == tmp2911 {
tmp2903 := PrimTail(V800)

tmp2904 := PrimTail(tmp2903)

tmp2905 := PrimTail(tmp2904)

tmp2906 := PrimTail(tmp2905)

tmp2907 := PrimEqual(Nil, tmp2906)

var ifres2902 Obj

if True == tmp2907 {
ifres2902 = True


} else {
ifres2902 = False


}

ifres2901 = ifres2902


} else {
ifres2901 = False


}

var ifres2900 Obj

if True == ifres2901 {
ifres2900 = True


} else {
ifres2900 = False


}

ifres2899 = ifres2900


} else {
ifres2899 = False


}

var ifres2898 Obj

if True == ifres2899 {
ifres2898 = True


} else {
ifres2898 = False


}

ifres2897 = ifres2898


} else {
ifres2897 = False


}

var ifres2896 Obj

if True == ifres2897 {
ifres2896 = True


} else {
ifres2896 = False


}

ifres2895 = ifres2896


} else {
ifres2895 = False


}

var ifres2894 Obj

if True == ifres2895 {
ifres2894 = True


} else {
ifres2894 = False


}

ifres2893 = ifres2894


} else {
ifres2893 = False


}

if True == ifres2893 {
tmp2871 := MakeNative(func(__e *ControlFlow) {
W804 := __e.Get(1)
_ = W804
tmp2872 := MakeNative(func(__e *ControlFlow) {
W805 := __e.Get(1)
_ = W805
tmp2873 := MakeNative(func(__e *ControlFlow) {
Z806 := __e.Get(1)
_ = Z806
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z806)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2873, W805)
return


}, 1)

tmp2874 := PrimTail(V800)

tmp2875 := PrimTail(tmp2874)

tmp2876 := PrimHead(tmp2875)

tmp2877 := PrimTail(V800)

tmp2878 := PrimHead(tmp2877)

tmp2879 := PrimTail(V800)

tmp2880 := PrimTail(tmp2879)

tmp2881 := PrimTail(tmp2880)

tmp2882 := PrimHead(tmp2881)

tmp2883 := Call(__e, PrimFunc(symshen_4beta), tmp2878, W804, tmp2882)


tmp2884 := PrimCons(tmp2883, Nil)

tmp2885 := PrimCons(tmp2876, tmp2884)

tmp2886 := PrimCons(W804, tmp2885)

tmp2887 := PrimCons(symlet, tmp2886)

__e.TailApply(tmp2872, tmp2887)
return


}, 1)

tmp2888 := Call(__e, PrimFunc(symgensym), symW)


__e.TailApply(tmp2871, tmp2888)
return


} else {
tmp2891 := PrimIsPair(V800)

if True == tmp2891 {
tmp2889 := MakeNative(func(__e *ControlFlow) {
Z807 := __e.Get(1)
_ = Z807
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z807)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2889, V800)
return


} else {
__e.Return(V800)
return
}


}


}


}, 1)

tmp2941 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp2857)


_ = tmp2941

tmp2942 := MakeNative(func(__e *ControlFlow) {
V808 := __e.Get(1)
_ = V808
V809 := __e.Get(2)
_ = V809
tmp2943 := MakeNative(func(__e *ControlFlow) {
Z810 := __e.Get(1)
_ = Z810
tmp2944 := Call(__e, PrimFunc(symfst), Z810)


tmp2945 := Call(__e, PrimFunc(symsnd), Z810)


tmp2946 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp2945)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp2944, V809, tmp2946)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp2943, V808)
return


}, 2)

tmp2947 := Call(__e, ns2_1set, symshen_4kl_1body, tmp2942)


_ = tmp2947

tmp2948 := MakeNative(func(__e *ControlFlow) {
V819 := __e.Get(1)
_ = V819
V820 := __e.Get(2)
_ = V820
V821 := __e.Get(3)
_ = V821
V822 := __e.Get(4)
_ = V822
tmp3097 := PrimEqual(Nil, V820)

var ifres3072 Obj

if True == tmp3097 {
tmp3096 := PrimEqual(Nil, V821)

var ifres3074 Obj

if True == tmp3096 {
tmp3095 := PrimIsPair(V822)

var ifres3076 Obj

if True == tmp3095 {
tmp3093 := PrimHead(V822)

tmp3094 := PrimEqual(symwhere, tmp3093)

var ifres3078 Obj

if True == tmp3094 {
tmp3091 := PrimTail(V822)

tmp3092 := PrimIsPair(tmp3091)

var ifres3080 Obj

if True == tmp3092 {
tmp3088 := PrimTail(V822)

tmp3089 := PrimTail(tmp3088)

tmp3090 := PrimIsPair(tmp3089)

var ifres3082 Obj

if True == tmp3090 {
tmp3084 := PrimTail(V822)

tmp3085 := PrimTail(tmp3084)

tmp3086 := PrimTail(tmp3085)

tmp3087 := PrimEqual(Nil, tmp3086)

var ifres3083 Obj

if True == tmp3087 {
ifres3083 = True


} else {
ifres3083 = False


}

ifres3082 = ifres3083


} else {
ifres3082 = False


}

var ifres3081 Obj

if True == ifres3082 {
ifres3081 = True


} else {
ifres3081 = False


}

ifres3080 = ifres3081


} else {
ifres3080 = False


}

var ifres3079 Obj

if True == ifres3080 {
ifres3079 = True


} else {
ifres3079 = False


}

ifres3078 = ifres3079


} else {
ifres3078 = False


}

var ifres3077 Obj

if True == ifres3078 {
ifres3077 = True


} else {
ifres3077 = False


}

ifres3076 = ifres3077


} else {
ifres3076 = False


}

var ifres3075 Obj

if True == ifres3076 {
ifres3075 = True


} else {
ifres3075 = False


}

ifres3074 = ifres3075


} else {
ifres3074 = False


}

var ifres3073 Obj

if True == ifres3074 {
ifres3073 = True


} else {
ifres3073 = False


}

ifres3072 = ifres3073


} else {
ifres3072 = False


}

if True == ifres3072 {
tmp2949 := PrimTail(V822)

tmp2950 := PrimHead(tmp2949)

tmp2951 := PrimCons(tmp2950, V819)

tmp2952 := PrimTail(V822)

tmp2953 := PrimTail(tmp2952)

tmp2954 := PrimHead(tmp2953)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp2951, Nil, Nil, tmp2954)
return


} else {
tmp3070 := PrimEqual(Nil, V820)

var ifres3067 Obj

if True == tmp3070 {
tmp3069 := PrimEqual(Nil, V821)

var ifres3068 Obj

if True == tmp3069 {
ifres3068 = True


} else {
ifres3068 = False


}

ifres3067 = ifres3068


} else {
ifres3067 = False


}

if True == ifres3067 {
tmp2955 := Call(__e, PrimFunc(symreverse), V819)


tmp2956 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp2955)


tmp2957 := PrimCons(V822, Nil)

__e.Return(PrimCons(tmp2956, tmp2957))
return


} else {
tmp3065 := PrimIsPair(V820)

var ifres3058 Obj

if True == tmp3065 {
tmp3064 := PrimIsPair(V821)

var ifres3060 Obj

if True == tmp3064 {
tmp3062 := PrimHead(V820)

tmp3063 := PrimIsVariable(tmp3062)

var ifres3061 Obj

if True == tmp3063 {
ifres3061 = True


} else {
ifres3061 = False


}

ifres3060 = ifres3061


} else {
ifres3060 = False


}

var ifres3059 Obj

if True == ifres3060 {
ifres3059 = True


} else {
ifres3059 = False


}

ifres3058 = ifres3059


} else {
ifres3058 = False


}

if True == ifres3058 {
tmp2958 := PrimTail(V820)

tmp2959 := PrimTail(V821)

tmp2960 := PrimHead(V820)

tmp2961 := PrimHead(V821)

tmp2962 := Call(__e, PrimFunc(symshen_4beta), tmp2960, tmp2961, V822)


__e.TailApply(PrimFunc(symshen_4triple_1stack), V819, tmp2958, tmp2959, tmp2962)
return


} else {
tmp3056 := PrimIsPair(V820)

var ifres3049 Obj

if True == tmp3056 {
tmp3055 := PrimIsPair(V821)

var ifres3051 Obj

if True == tmp3055 {
tmp3053 := PrimHead(V820)

tmp3054 := Call(__e, PrimFunc(symshen_4custom_1pattern_2), tmp3053)


var ifres3052 Obj

if True == tmp3054 {
ifres3052 = True


} else {
ifres3052 = False


}

ifres3051 = ifres3052


} else {
ifres3051 = False


}

var ifres3050 Obj

if True == ifres3051 {
ifres3050 = True


} else {
ifres3050 = False


}

ifres3049 = ifres3050


} else {
ifres3049 = False


}

if True == ifres3049 {
tmp2963 := PrimHead(V820)

tmp2964 := Call(__e, PrimFunc(symshen_4custom_1pattern_1body), tmp2963)


tmp2965 := PrimTail(V820)

tmp2966 := PrimHead(V821)

tmp2967 := PrimTail(V821)

__e.TailApply(PrimFunc(symshen_4custom_1pattern_1triple_1stack), V819, tmp2964, tmp2965, tmp2966, tmp2967, V822)
return


} else {
tmp3047 := PrimIsPair(V820)

var ifres3017 Obj

if True == tmp3047 {
tmp3045 := PrimHead(V820)

tmp3046 := PrimIsPair(tmp3045)

var ifres3019 Obj

if True == tmp3046 {
tmp3042 := PrimHead(V820)

tmp3043 := PrimTail(tmp3042)

tmp3044 := PrimIsPair(tmp3043)

var ifres3021 Obj

if True == tmp3044 {
tmp3038 := PrimHead(V820)

tmp3039 := PrimTail(tmp3038)

tmp3040 := PrimTail(tmp3039)

tmp3041 := PrimIsPair(tmp3040)

var ifres3023 Obj

if True == tmp3041 {
tmp3033 := PrimHead(V820)

tmp3034 := PrimTail(tmp3033)

tmp3035 := PrimTail(tmp3034)

tmp3036 := PrimTail(tmp3035)

tmp3037 := PrimEqual(Nil, tmp3036)

var ifres3025 Obj

if True == tmp3037 {
tmp3032 := PrimIsPair(V821)

var ifres3027 Obj

if True == tmp3032 {
tmp3029 := PrimHead(V820)

tmp3030 := PrimHead(tmp3029)

tmp3031 := Call(__e, PrimFunc(symshen_4constructor_2), tmp3030)


var ifres3028 Obj

if True == tmp3031 {
ifres3028 = True


} else {
ifres3028 = False


}

ifres3027 = ifres3028


} else {
ifres3027 = False


}

var ifres3026 Obj

if True == ifres3027 {
ifres3026 = True


} else {
ifres3026 = False


}

ifres3025 = ifres3026


} else {
ifres3025 = False


}

var ifres3024 Obj

if True == ifres3025 {
ifres3024 = True


} else {
ifres3024 = False


}

ifres3023 = ifres3024


} else {
ifres3023 = False


}

var ifres3022 Obj

if True == ifres3023 {
ifres3022 = True


} else {
ifres3022 = False


}

ifres3021 = ifres3022


} else {
ifres3021 = False


}

var ifres3020 Obj

if True == ifres3021 {
ifres3020 = True


} else {
ifres3020 = False


}

ifres3019 = ifres3020


} else {
ifres3019 = False


}

var ifres3018 Obj

if True == ifres3019 {
ifres3018 = True


} else {
ifres3018 = False


}

ifres3017 = ifres3018


} else {
ifres3017 = False


}

if True == ifres3017 {
tmp2968 := PrimHead(V820)

tmp2969 := PrimHead(tmp2968)

tmp2970 := Call(__e, PrimFunc(symshen_4op_1test), tmp2969)


tmp2971 := PrimHead(V821)

tmp2972 := PrimCons(tmp2971, Nil)

tmp2973 := PrimCons(tmp2970, tmp2972)

tmp2974 := PrimCons(tmp2973, V819)

tmp2975 := PrimHead(V820)

tmp2976 := PrimTail(tmp2975)

tmp2977 := PrimHead(tmp2976)

tmp2978 := PrimHead(V820)

tmp2979 := PrimTail(tmp2978)

tmp2980 := PrimTail(tmp2979)

tmp2981 := PrimHead(tmp2980)

tmp2982 := PrimTail(V820)

tmp2983 := PrimCons(tmp2981, tmp2982)

tmp2984 := PrimCons(tmp2977, tmp2983)

tmp2985 := PrimHead(V820)

tmp2986 := PrimHead(tmp2985)

tmp2987 := Call(__e, PrimFunc(symshen_4op1), tmp2986)


tmp2988 := PrimHead(V821)

tmp2989 := PrimCons(tmp2988, Nil)

tmp2990 := PrimCons(tmp2987, tmp2989)

tmp2991 := PrimHead(V820)

tmp2992 := PrimHead(tmp2991)

tmp2993 := Call(__e, PrimFunc(symshen_4op2), tmp2992)


tmp2994 := PrimHead(V821)

tmp2995 := PrimCons(tmp2994, Nil)

tmp2996 := PrimCons(tmp2993, tmp2995)

tmp2997 := PrimTail(V821)

tmp2998 := PrimCons(tmp2996, tmp2997)

tmp2999 := PrimCons(tmp2990, tmp2998)

tmp3000 := PrimHead(V820)

tmp3001 := PrimHead(V821)

tmp3002 := Call(__e, PrimFunc(symshen_4beta), tmp3000, tmp3001, V822)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp2974, tmp2984, tmp2999, tmp3002)
return


} else {
tmp3015 := PrimIsPair(V820)

var ifres3012 Obj

if True == tmp3015 {
tmp3014 := PrimIsPair(V821)

var ifres3013 Obj

if True == tmp3014 {
ifres3013 = True


} else {
ifres3013 = False


}

ifres3012 = ifres3013


} else {
ifres3012 = False


}

if True == ifres3012 {
tmp3003 := PrimHead(V820)

tmp3004 := PrimHead(V821)

tmp3005 := PrimCons(tmp3004, Nil)

tmp3006 := PrimCons(tmp3003, tmp3005)

tmp3007 := PrimCons(sym_a, tmp3006)

tmp3008 := PrimCons(tmp3007, V819)

tmp3009 := PrimTail(V820)

tmp3010 := PrimTail(V821)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3008, tmp3009, tmp3010, V822)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.triple-stack")))
return
}


}


}


}


}


}


}, 4)

tmp3098 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp2948)


_ = tmp3098

tmp3099 := MakeNative(func(__e *ControlFlow) {
V823 := __e.Get(1)
_ = V823
V824 := __e.Get(2)
_ = V824
V825 := __e.Get(3)
_ = V825
V826 := __e.Get(4)
_ = V826
V827 := __e.Get(5)
_ = V827
V828 := __e.Get(6)
_ = V828
tmp3100 := MakeNative(func(__e *ControlFlow) {
W829 := __e.Get(1)
_ = W829
tmp3115 := Call(__e, PrimFunc(symtuple_2), W829)


if True == tmp3115 {
tmp3101 := MakeNative(func(__e *ControlFlow) {
W830 := __e.Get(1)
_ = W830
tmp3102 := MakeNative(func(__e *ControlFlow) {
W831 := __e.Get(1)
_ = W831
tmp3103 := Call(__e, PrimFunc(symreverse), W830)


tmp3104 := Call(__e, PrimFunc(symappend), tmp3103, V823)


tmp3105 := MakeNative(func(__e *ControlFlow) {
Z832 := __e.Get(1)
_ = Z832
__e.TailApply(PrimFunc(symfst), Z832)
return
}, 1)

tmp3106 := Call(__e, PrimFunc(symmap), tmp3105, W831)


tmp3107 := Call(__e, PrimFunc(symappend), tmp3106, V825)


tmp3108 := MakeNative(func(__e *ControlFlow) {
Z833 := __e.Get(1)
_ = Z833
__e.TailApply(PrimFunc(symsnd), Z833)
return
}, 1)

tmp3109 := Call(__e, PrimFunc(symmap), tmp3108, W831)


tmp3110 := Call(__e, PrimFunc(symappend), tmp3109, V827)


tmp3111 := Call(__e, PrimFunc(symshen_4beta), V824, V826, V828)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3104, tmp3107, tmp3110, tmp3111)
return


}, 1)

tmp3112 := Call(__e, PrimFunc(symsnd), W829)


__e.TailApply(tmp3102, tmp3112)
return


}, 1)

tmp3113 := Call(__e, PrimFunc(symfst), W829)


__e.TailApply(tmp3101, tmp3113)
return


} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V824)
return
}


}, 1)

tmp3116 := Call(__e, PrimFunc(sym_8p), V824, V826)


tmp3117 := Call(__e, PrimFunc(symshen_4custom_1pattern_1reducer), tmp3116)


__e.TailApply(tmp3100, tmp3117)
return


}, 6)

tmp3118 := Call(__e, ns2_1set, symshen_4custom_1pattern_1triple_1stack, tmp3099)


_ = tmp3118

tmp3119 := MakeNative(func(__e *ControlFlow) {
V836 := __e.Get(1)
_ = V836
tmp3138 := PrimEqual(Nil, V836)

if True == tmp3138 {
__e.Return(True)
return
} else {
tmp3136 := PrimIsPair(V836)

var ifres3132 Obj

if True == tmp3136 {
tmp3134 := PrimTail(V836)

tmp3135 := PrimEqual(Nil, tmp3134)

var ifres3133 Obj

if True == tmp3135 {
ifres3133 = True


} else {
ifres3133 = False


}

ifres3132 = ifres3133


} else {
ifres3132 = False


}

if True == ifres3132 {
__e.Return(PrimHead(V836))
return
} else {
tmp3130 := PrimIsPair(V836)

var ifres3126 Obj

if True == tmp3130 {
tmp3128 := PrimTail(V836)

tmp3129 := PrimIsPair(tmp3128)

var ifres3127 Obj

if True == tmp3129 {
ifres3127 = True


} else {
ifres3127 = False


}

ifres3126 = ifres3127


} else {
ifres3126 = False


}

if True == ifres3126 {
tmp3120 := PrimHead(V836)

tmp3121 := PrimTail(V836)

tmp3122 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp3121)


tmp3123 := PrimCons(tmp3122, Nil)

tmp3124 := PrimCons(tmp3120, tmp3123)

__e.Return(PrimCons(symand, tmp3124))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rectify-test")))
return
}


}


}


}, 1)

tmp3139 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp3119)


_ = tmp3139

tmp3140 := MakeNative(func(__e *ControlFlow) {
V846 := __e.Get(1)
_ = V846
V847 := __e.Get(2)
_ = V847
V848 := __e.Get(3)
_ = V848
tmp3217 := PrimEqual(V846, V848)

if True == tmp3217 {
__e.Return(V847)
return
} else {
tmp3215 := PrimIsPair(V848)

var ifres3191 Obj

if True == tmp3215 {
tmp3213 := PrimHead(V848)

tmp3214 := PrimEqual(symlambda, tmp3213)

var ifres3193 Obj

if True == tmp3214 {
tmp3211 := PrimTail(V848)

tmp3212 := PrimIsPair(tmp3211)

var ifres3195 Obj

if True == tmp3212 {
tmp3208 := PrimTail(V848)

tmp3209 := PrimTail(tmp3208)

tmp3210 := PrimIsPair(tmp3209)

var ifres3197 Obj

if True == tmp3210 {
tmp3204 := PrimTail(V848)

tmp3205 := PrimTail(tmp3204)

tmp3206 := PrimTail(tmp3205)

tmp3207 := PrimEqual(Nil, tmp3206)

var ifres3199 Obj

if True == tmp3207 {
tmp3201 := PrimTail(V848)

tmp3202 := PrimHead(tmp3201)

tmp3203 := PrimEqual(V846, tmp3202)

var ifres3200 Obj

if True == tmp3203 {
ifres3200 = True


} else {
ifres3200 = False


}

ifres3199 = ifres3200


} else {
ifres3199 = False


}

var ifres3198 Obj

if True == ifres3199 {
ifres3198 = True


} else {
ifres3198 = False


}

ifres3197 = ifres3198


} else {
ifres3197 = False


}

var ifres3196 Obj

if True == ifres3197 {
ifres3196 = True


} else {
ifres3196 = False


}

ifres3195 = ifres3196


} else {
ifres3195 = False


}

var ifres3194 Obj

if True == ifres3195 {
ifres3194 = True


} else {
ifres3194 = False


}

ifres3193 = ifres3194


} else {
ifres3193 = False


}

var ifres3192 Obj

if True == ifres3193 {
ifres3192 = True


} else {
ifres3192 = False


}

ifres3191 = ifres3192


} else {
ifres3191 = False


}

if True == ifres3191 {
__e.Return(V848)
return
} else {
tmp3189 := PrimIsPair(V848)

var ifres3158 Obj

if True == tmp3189 {
tmp3187 := PrimHead(V848)

tmp3188 := PrimEqual(symlet, tmp3187)

var ifres3160 Obj

if True == tmp3188 {
tmp3185 := PrimTail(V848)

tmp3186 := PrimIsPair(tmp3185)

var ifres3162 Obj

if True == tmp3186 {
tmp3182 := PrimTail(V848)

tmp3183 := PrimTail(tmp3182)

tmp3184 := PrimIsPair(tmp3183)

var ifres3164 Obj

if True == tmp3184 {
tmp3178 := PrimTail(V848)

tmp3179 := PrimTail(tmp3178)

tmp3180 := PrimTail(tmp3179)

tmp3181 := PrimIsPair(tmp3180)

var ifres3166 Obj

if True == tmp3181 {
tmp3173 := PrimTail(V848)

tmp3174 := PrimTail(tmp3173)

tmp3175 := PrimTail(tmp3174)

tmp3176 := PrimTail(tmp3175)

tmp3177 := PrimEqual(Nil, tmp3176)

var ifres3168 Obj

if True == tmp3177 {
tmp3170 := PrimTail(V848)

tmp3171 := PrimHead(tmp3170)

tmp3172 := PrimEqual(V846, tmp3171)

var ifres3169 Obj

if True == tmp3172 {
ifres3169 = True


} else {
ifres3169 = False


}

ifres3168 = ifres3169


} else {
ifres3168 = False


}

var ifres3167 Obj

if True == ifres3168 {
ifres3167 = True


} else {
ifres3167 = False


}

ifres3166 = ifres3167


} else {
ifres3166 = False


}

var ifres3165 Obj

if True == ifres3166 {
ifres3165 = True


} else {
ifres3165 = False


}

ifres3164 = ifres3165


} else {
ifres3164 = False


}

var ifres3163 Obj

if True == ifres3164 {
ifres3163 = True


} else {
ifres3163 = False


}

ifres3162 = ifres3163


} else {
ifres3162 = False


}

var ifres3161 Obj

if True == ifres3162 {
ifres3161 = True


} else {
ifres3161 = False


}

ifres3160 = ifres3161


} else {
ifres3160 = False


}

var ifres3159 Obj

if True == ifres3160 {
ifres3159 = True


} else {
ifres3159 = False


}

ifres3158 = ifres3159


} else {
ifres3158 = False


}

if True == ifres3158 {
tmp3141 := PrimTail(V848)

tmp3142 := PrimHead(tmp3141)

tmp3143 := PrimTail(V848)

tmp3144 := PrimHead(tmp3143)

tmp3145 := PrimTail(V848)

tmp3146 := PrimTail(tmp3145)

tmp3147 := PrimHead(tmp3146)

tmp3148 := Call(__e, PrimFunc(symshen_4beta), tmp3144, V847, tmp3147)


tmp3149 := PrimTail(V848)

tmp3150 := PrimTail(tmp3149)

tmp3151 := PrimTail(tmp3150)

tmp3152 := PrimCons(tmp3148, tmp3151)

tmp3153 := PrimCons(tmp3142, tmp3152)

__e.Return(PrimCons(symlet, tmp3153))
return


} else {
tmp3156 := PrimIsPair(V848)

if True == tmp3156 {
tmp3154 := MakeNative(func(__e *ControlFlow) {
Z849 := __e.Get(1)
_ = Z849
__e.TailApply(PrimFunc(symshen_4beta), V846, V847, Z849)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3154, V848)
return


} else {
__e.Return(V848)
return
}


}


}


}


}, 3)

tmp3218 := Call(__e, ns2_1set, symshen_4beta, tmp3140)


_ = tmp3218

tmp3219 := MakeNative(func(__e *ControlFlow) {
V852 := __e.Get(1)
_ = V852
tmp3227 := PrimEqual(symcons, V852)

if True == tmp3227 {
__e.Return(symhd)
return
} else {
tmp3225 := PrimEqual(sym_8s, V852)

if True == tmp3225 {
__e.Return(symhdstr)
return
} else {
tmp3223 := PrimEqual(sym_8p, V852)

if True == tmp3223 {
__e.Return(symfst)
return
} else {
tmp3221 := PrimEqual(sym_8v, V852)

if True == tmp3221 {
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

tmp3228 := Call(__e, ns2_1set, symshen_4op1, tmp3219)


_ = tmp3228

tmp3229 := MakeNative(func(__e *ControlFlow) {
V855 := __e.Get(1)
_ = V855
tmp3237 := PrimEqual(symcons, V855)

if True == tmp3237 {
__e.Return(symtl)
return
} else {
tmp3235 := PrimEqual(sym_8s, V855)

if True == tmp3235 {
__e.Return(symtlstr)
return
} else {
tmp3233 := PrimEqual(sym_8p, V855)

if True == tmp3233 {
__e.Return(symsnd)
return
} else {
tmp3231 := PrimEqual(sym_8v, V855)

if True == tmp3231 {
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

tmp3238 := Call(__e, ns2_1set, symshen_4op2, tmp3229)


_ = tmp3238

tmp3239 := MakeNative(func(__e *ControlFlow) {
V858 := __e.Get(1)
_ = V858
tmp3247 := PrimEqual(symcons, V858)

if True == tmp3247 {
__e.Return(symcons_2)
return
} else {
tmp3245 := PrimEqual(sym_8s, V858)

if True == tmp3245 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp3243 := PrimEqual(sym_8p, V858)

if True == tmp3243 {
__e.Return(symtuple_2)
return
} else {
tmp3241 := PrimEqual(sym_8v, V858)

if True == tmp3241 {
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

tmp3248 := Call(__e, ns2_1set, symshen_4op_1test, tmp3239)


_ = tmp3248

tmp3249 := MakeNative(func(__e *ControlFlow) {
V859 := __e.Get(1)
_ = V859
tmp3251 := PrimEqual(MakeString(""), V859)

if True == tmp3251 {
__e.Return(False)
return
} else {
__e.Return(PrimIsString(V859))
return
}


}, 1)

tmp3252 := Call(__e, ns2_1set, symshen_4_7string_2, tmp3249)


_ = tmp3252

tmp3253 := MakeNative(func(__e *ControlFlow) {
V860 := __e.Get(1)
_ = V860
tmp3255 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp3256 := PrimEqual(V860, tmp3255)

if True == tmp3256 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V860)
return
}


}, 1)

tmp3257 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp3253)


_ = tmp3257

tmp3258 := MakeNative(func(__e *ControlFlow) {
V863 := __e.Get(1)
_ = V863
tmp3262 := PrimEqual(sym_7, V863)

if True == tmp3262 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, True))
return
} else {
tmp3260 := PrimEqual(sym_1, V863)

if True == tmp3260 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("factorise expects a + or a -\n")))
return
}


}


}, 1)

tmp3263 := Call(__e, ns2_1set, symfactorise, tmp3258)


_ = tmp3263

tmp3264 := MakeNative(func(__e *ControlFlow) {
V864 := __e.Get(1)
_ = V864
tmp3266 := PrimValue(symshen_4_dfactorise_2_d)

if True == tmp3266 {
__e.TailApply(PrimFunc(symshen_4factor), V864)
return
} else {
__e.Return(V864)
return
}


}, 1)

tmp3267 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp3264)


_ = tmp3267

tmp3268 := MakeNative(func(__e *ControlFlow) {
V865 := __e.Get(1)
_ = V865
tmp3325 := PrimIsPair(V865)

var ifres3284 Obj

if True == tmp3325 {
tmp3323 := PrimHead(V865)

tmp3324 := PrimEqual(symdefun, tmp3323)

var ifres3286 Obj

if True == tmp3324 {
tmp3321 := PrimTail(V865)

tmp3322 := PrimIsPair(tmp3321)

var ifres3288 Obj

if True == tmp3322 {
tmp3318 := PrimTail(V865)

tmp3319 := PrimTail(tmp3318)

tmp3320 := PrimIsPair(tmp3319)

var ifres3290 Obj

if True == tmp3320 {
tmp3314 := PrimTail(V865)

tmp3315 := PrimTail(tmp3314)

tmp3316 := PrimTail(tmp3315)

tmp3317 := PrimIsPair(tmp3316)

var ifres3292 Obj

if True == tmp3317 {
tmp3309 := PrimTail(V865)

tmp3310 := PrimTail(tmp3309)

tmp3311 := PrimTail(tmp3310)

tmp3312 := PrimHead(tmp3311)

tmp3313 := PrimIsPair(tmp3312)

var ifres3294 Obj

if True == tmp3313 {
tmp3303 := PrimTail(V865)

tmp3304 := PrimTail(tmp3303)

tmp3305 := PrimTail(tmp3304)

tmp3306 := PrimHead(tmp3305)

tmp3307 := PrimHead(tmp3306)

tmp3308 := PrimEqual(symcond, tmp3307)

var ifres3296 Obj

if True == tmp3308 {
tmp3298 := PrimTail(V865)

tmp3299 := PrimTail(tmp3298)

tmp3300 := PrimTail(tmp3299)

tmp3301 := PrimTail(tmp3300)

tmp3302 := PrimEqual(Nil, tmp3301)

var ifres3297 Obj

if True == tmp3302 {
ifres3297 = True


} else {
ifres3297 = False


}

ifres3296 = ifres3297


} else {
ifres3296 = False


}

var ifres3295 Obj

if True == ifres3296 {
ifres3295 = True


} else {
ifres3295 = False


}

ifres3294 = ifres3295


} else {
ifres3294 = False


}

var ifres3293 Obj

if True == ifres3294 {
ifres3293 = True


} else {
ifres3293 = False


}

ifres3292 = ifres3293


} else {
ifres3292 = False


}

var ifres3291 Obj

if True == ifres3292 {
ifres3291 = True


} else {
ifres3291 = False


}

ifres3290 = ifres3291


} else {
ifres3290 = False


}

var ifres3289 Obj

if True == ifres3290 {
ifres3289 = True


} else {
ifres3289 = False


}

ifres3288 = ifres3289


} else {
ifres3288 = False


}

var ifres3287 Obj

if True == ifres3288 {
ifres3287 = True


} else {
ifres3287 = False


}

ifres3286 = ifres3287


} else {
ifres3286 = False


}

var ifres3285 Obj

if True == ifres3286 {
ifres3285 = True


} else {
ifres3285 = False


}

ifres3284 = ifres3285


} else {
ifres3284 = False


}

if True == ifres3284 {
tmp3269 := PrimTail(V865)

tmp3270 := PrimHead(tmp3269)

tmp3271 := PrimTail(V865)

tmp3272 := PrimTail(tmp3271)

tmp3273 := PrimHead(tmp3272)

tmp3274 := PrimTail(V865)

tmp3275 := PrimTail(tmp3274)

tmp3276 := PrimTail(tmp3275)

tmp3277 := PrimHead(tmp3276)

tmp3278 := PrimTail(tmp3277)

tmp3279 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp3278)


tmp3280 := PrimCons(tmp3279, Nil)

tmp3281 := PrimCons(tmp3273, tmp3280)

tmp3282 := PrimCons(tmp3270, tmp3281)

__e.Return(PrimCons(symdefun, tmp3282))
return


} else {
__e.Return(V865)
return
}


}, 1)

tmp3326 := Call(__e, ns2_1set, symshen_4factor, tmp3268)


_ = tmp3326

tmp3327 := MakeNative(func(__e *ControlFlow) {
V868 := __e.Get(1)
_ = V868
tmp3483 := PrimIsPair(V868)

var ifres3463 Obj

if True == tmp3483 {
tmp3481 := PrimHead(V868)

tmp3482 := PrimIsPair(tmp3481)

var ifres3465 Obj

if True == tmp3482 {
tmp3478 := PrimHead(V868)

tmp3479 := PrimHead(tmp3478)

tmp3480 := PrimEqual(True, tmp3479)

var ifres3467 Obj

if True == tmp3480 {
tmp3475 := PrimHead(V868)

tmp3476 := PrimTail(tmp3475)

tmp3477 := PrimIsPair(tmp3476)

var ifres3469 Obj

if True == tmp3477 {
tmp3471 := PrimHead(V868)

tmp3472 := PrimTail(tmp3471)

tmp3473 := PrimTail(tmp3472)

tmp3474 := PrimEqual(Nil, tmp3473)

var ifres3470 Obj

if True == tmp3474 {
ifres3470 = True


} else {
ifres3470 = False


}

ifres3469 = ifres3470


} else {
ifres3469 = False


}

var ifres3468 Obj

if True == ifres3469 {
ifres3468 = True


} else {
ifres3468 = False


}

ifres3467 = ifres3468


} else {
ifres3467 = False


}

var ifres3466 Obj

if True == ifres3467 {
ifres3466 = True


} else {
ifres3466 = False


}

ifres3465 = ifres3466


} else {
ifres3465 = False


}

var ifres3464 Obj

if True == ifres3465 {
ifres3464 = True


} else {
ifres3464 = False


}

ifres3463 = ifres3464


} else {
ifres3463 = False


}

if True == ifres3463 {
tmp3328 := PrimHead(V868)

tmp3329 := PrimTail(tmp3328)

__e.Return(PrimHead(tmp3329))
return


} else {
tmp3461 := PrimIsPair(V868)

var ifres3414 Obj

if True == tmp3461 {
tmp3459 := PrimHead(V868)

tmp3460 := PrimIsPair(tmp3459)

var ifres3416 Obj

if True == tmp3460 {
tmp3456 := PrimHead(V868)

tmp3457 := PrimHead(tmp3456)

tmp3458 := PrimIsPair(tmp3457)

var ifres3418 Obj

if True == tmp3458 {
tmp3452 := PrimHead(V868)

tmp3453 := PrimHead(tmp3452)

tmp3454 := PrimHead(tmp3453)

tmp3455 := PrimEqual(symand, tmp3454)

var ifres3420 Obj

if True == tmp3455 {
tmp3448 := PrimHead(V868)

tmp3449 := PrimHead(tmp3448)

tmp3450 := PrimTail(tmp3449)

tmp3451 := PrimIsPair(tmp3450)

var ifres3422 Obj

if True == tmp3451 {
tmp3443 := PrimHead(V868)

tmp3444 := PrimHead(tmp3443)

tmp3445 := PrimTail(tmp3444)

tmp3446 := PrimTail(tmp3445)

tmp3447 := PrimIsPair(tmp3446)

var ifres3424 Obj

if True == tmp3447 {
tmp3437 := PrimHead(V868)

tmp3438 := PrimHead(tmp3437)

tmp3439 := PrimTail(tmp3438)

tmp3440 := PrimTail(tmp3439)

tmp3441 := PrimTail(tmp3440)

tmp3442 := PrimEqual(Nil, tmp3441)

var ifres3426 Obj

if True == tmp3442 {
tmp3434 := PrimHead(V868)

tmp3435 := PrimTail(tmp3434)

tmp3436 := PrimIsPair(tmp3435)

var ifres3428 Obj

if True == tmp3436 {
tmp3430 := PrimHead(V868)

tmp3431 := PrimTail(tmp3430)

tmp3432 := PrimTail(tmp3431)

tmp3433 := PrimEqual(Nil, tmp3432)

var ifres3429 Obj

if True == tmp3433 {
ifres3429 = True


} else {
ifres3429 = False


}

ifres3428 = ifres3429


} else {
ifres3428 = False


}

var ifres3427 Obj

if True == ifres3428 {
ifres3427 = True


} else {
ifres3427 = False


}

ifres3426 = ifres3427


} else {
ifres3426 = False


}

var ifres3425 Obj

if True == ifres3426 {
ifres3425 = True


} else {
ifres3425 = False


}

ifres3424 = ifres3425


} else {
ifres3424 = False


}

var ifres3423 Obj

if True == ifres3424 {
ifres3423 = True


} else {
ifres3423 = False


}

ifres3422 = ifres3423


} else {
ifres3422 = False


}

var ifres3421 Obj

if True == ifres3422 {
ifres3421 = True


} else {
ifres3421 = False


}

ifres3420 = ifres3421


} else {
ifres3420 = False


}

var ifres3419 Obj

if True == ifres3420 {
ifres3419 = True


} else {
ifres3419 = False


}

ifres3418 = ifres3419


} else {
ifres3418 = False


}

var ifres3417 Obj

if True == ifres3418 {
ifres3417 = True


} else {
ifres3417 = False


}

ifres3416 = ifres3417


} else {
ifres3416 = False


}

var ifres3415 Obj

if True == ifres3416 {
ifres3415 = True


} else {
ifres3415 = False


}

ifres3414 = ifres3415


} else {
ifres3414 = False


}

if True == ifres3414 {
tmp3330 := MakeNative(func(__e *ControlFlow) {
W869 := __e.Get(1)
_ = W869
tmp3331 := MakeNative(func(__e *ControlFlow) {
W870 := __e.Get(1)
_ = W870
tmp3379 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), W870)


if True == tmp3379 {
tmp3332 := PrimHead(V868)

tmp3333 := PrimHead(tmp3332)

tmp3334 := PrimHead(V868)

tmp3335 := PrimTail(tmp3334)

tmp3336 := PrimHead(tmp3335)

tmp3337 := PrimTail(V868)

tmp3338 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp3337)


tmp3339 := PrimCons(tmp3338, Nil)

tmp3340 := PrimCons(tmp3336, tmp3339)

tmp3341 := PrimCons(tmp3333, tmp3340)

__e.Return(PrimCons(symif, tmp3341))
return


} else {
tmp3342 := MakeNative(func(__e *ControlFlow) {
W871 := __e.Get(1)
_ = W871
tmp3343 := MakeNative(func(__e *ControlFlow) {
W872 := __e.Get(1)
_ = W872
tmp3344 := MakeNative(func(__e *ControlFlow) {
W873 := __e.Get(1)
_ = W873
tmp3345 := MakeNative(func(__e *ControlFlow) {
W874 := __e.Get(1)
_ = W874
tmp3346 := MakeNative(func(__e *ControlFlow) {
W875 := __e.Get(1)
_ = W875
__e.TailApply(PrimFunc(symshen_4remove_1indirection), W875)
return
}, 1)

tmp3347 := PrimCons(W872, Nil)

tmp3348 := PrimCons(symfreeze, tmp3347)

tmp3349 := PrimHead(V868)

tmp3350 := PrimHead(tmp3349)

tmp3351 := PrimTail(tmp3350)

tmp3352 := PrimHead(tmp3351)

tmp3353 := PrimHead(V868)

tmp3354 := PrimHead(tmp3353)

tmp3355 := PrimTail(tmp3354)

tmp3356 := PrimHead(tmp3355)

tmp3357 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W874)


tmp3358 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp3356, tmp3357)


tmp3359 := PrimCons(W873, Nil)

tmp3360 := PrimCons(symthaw, tmp3359)

tmp3361 := PrimCons(tmp3360, Nil)

tmp3362 := PrimCons(tmp3358, tmp3361)

tmp3363 := PrimCons(tmp3352, tmp3362)

tmp3364 := PrimCons(symif, tmp3363)

tmp3365 := PrimCons(tmp3364, Nil)

tmp3366 := PrimCons(tmp3348, tmp3365)

tmp3367 := PrimCons(W873, tmp3366)

tmp3368 := PrimCons(symlet, tmp3367)

__e.TailApply(tmp3346, tmp3368)
return


}, 1)

tmp3369 := PrimCons(W873, Nil)

tmp3370 := PrimCons(symthaw, tmp3369)

tmp3371 := PrimCons(tmp3370, Nil)

tmp3372 := PrimCons(True, tmp3371)

tmp3373 := PrimCons(tmp3372, W870)

tmp3374 := Call(__e, PrimFunc(symreverse), tmp3373)


__e.TailApply(tmp3345, tmp3374)
return


}, 1)

tmp3375 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp3344, tmp3375)
return


}, 1)

tmp3376 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W871)


__e.TailApply(tmp3343, tmp3376)
return


}, 1)

tmp3377 := Call(__e, PrimFunc(symsnd), W869)


__e.TailApply(tmp3342, tmp3377)
return


}


}, 1)

tmp3380 := Call(__e, PrimFunc(symfst), W869)


__e.TailApply(tmp3331, tmp3380)
return


}, 1)

tmp3381 := PrimHead(V868)

tmp3382 := PrimHead(tmp3381)

tmp3383 := PrimTail(tmp3382)

tmp3384 := PrimHead(tmp3383)

tmp3385 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp3384, V868, Nil)


__e.TailApply(tmp3330, tmp3385)
return


} else {
tmp3412 := PrimIsPair(V868)

var ifres3397 Obj

if True == tmp3412 {
tmp3410 := PrimHead(V868)

tmp3411 := PrimIsPair(tmp3410)

var ifres3399 Obj

if True == tmp3411 {
tmp3407 := PrimHead(V868)

tmp3408 := PrimTail(tmp3407)

tmp3409 := PrimIsPair(tmp3408)

var ifres3401 Obj

if True == tmp3409 {
tmp3403 := PrimHead(V868)

tmp3404 := PrimTail(tmp3403)

tmp3405 := PrimTail(tmp3404)

tmp3406 := PrimEqual(Nil, tmp3405)

var ifres3402 Obj

if True == tmp3406 {
ifres3402 = True


} else {
ifres3402 = False


}

ifres3401 = ifres3402


} else {
ifres3401 = False


}

var ifres3400 Obj

if True == ifres3401 {
ifres3400 = True


} else {
ifres3400 = False


}

ifres3399 = ifres3400


} else {
ifres3399 = False


}

var ifres3398 Obj

if True == ifres3399 {
ifres3398 = True


} else {
ifres3398 = False


}

ifres3397 = ifres3398


} else {
ifres3397 = False


}

if True == ifres3397 {
tmp3386 := PrimHead(V868)

tmp3387 := PrimHead(tmp3386)

tmp3388 := PrimHead(V868)

tmp3389 := PrimTail(tmp3388)

tmp3390 := PrimHead(tmp3389)

tmp3391 := PrimTail(V868)

tmp3392 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp3391)


tmp3393 := PrimCons(tmp3392, Nil)

tmp3394 := PrimCons(tmp3390, tmp3393)

tmp3395 := PrimCons(tmp3387, tmp3394)

__e.Return(PrimCons(symif, tmp3395))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1recognisors)
return
}


}


}


}, 1)

tmp3484 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp3327)


_ = tmp3484

tmp3485 := MakeNative(func(__e *ControlFlow) {
V880 := __e.Get(1)
_ = V880
tmp3491 := PrimIsPair(V880)

var ifres3487 Obj

if True == tmp3491 {
tmp3489 := PrimTail(V880)

tmp3490 := PrimEqual(Nil, tmp3489)

var ifres3488 Obj

if True == tmp3490 {
ifres3488 = True


} else {
ifres3488 = False


}

ifres3487 = ifres3488


} else {
ifres3487 = False


}

if True == ifres3487 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3492 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp3485)


_ = tmp3492

tmp3493 := MakeNative(func(__e *ControlFlow) {
V881 := __e.Get(1)
_ = V881
tmp3608 := PrimIsPair(V881)

var ifres3508 Obj

if True == tmp3608 {
tmp3606 := PrimHead(V881)

tmp3607 := PrimEqual(symlet, tmp3606)

var ifres3510 Obj

if True == tmp3607 {
tmp3604 := PrimTail(V881)

tmp3605 := PrimIsPair(tmp3604)

var ifres3512 Obj

if True == tmp3605 {
tmp3601 := PrimTail(V881)

tmp3602 := PrimTail(tmp3601)

tmp3603 := PrimIsPair(tmp3602)

var ifres3514 Obj

if True == tmp3603 {
tmp3597 := PrimTail(V881)

tmp3598 := PrimTail(tmp3597)

tmp3599 := PrimHead(tmp3598)

tmp3600 := PrimIsPair(tmp3599)

var ifres3516 Obj

if True == tmp3600 {
tmp3592 := PrimTail(V881)

tmp3593 := PrimTail(tmp3592)

tmp3594 := PrimHead(tmp3593)

tmp3595 := PrimHead(tmp3594)

tmp3596 := PrimEqual(symfreeze, tmp3595)

var ifres3518 Obj

if True == tmp3596 {
tmp3587 := PrimTail(V881)

tmp3588 := PrimTail(tmp3587)

tmp3589 := PrimHead(tmp3588)

tmp3590 := PrimTail(tmp3589)

tmp3591 := PrimIsPair(tmp3590)

var ifres3520 Obj

if True == tmp3591 {
tmp3581 := PrimTail(V881)

tmp3582 := PrimTail(tmp3581)

tmp3583 := PrimHead(tmp3582)

tmp3584 := PrimTail(tmp3583)

tmp3585 := PrimHead(tmp3584)

tmp3586 := PrimIsPair(tmp3585)

var ifres3522 Obj

if True == tmp3586 {
tmp3574 := PrimTail(V881)

tmp3575 := PrimTail(tmp3574)

tmp3576 := PrimHead(tmp3575)

tmp3577 := PrimTail(tmp3576)

tmp3578 := PrimHead(tmp3577)

tmp3579 := PrimHead(tmp3578)

tmp3580 := PrimEqual(symthaw, tmp3579)

var ifres3524 Obj

if True == tmp3580 {
tmp3567 := PrimTail(V881)

tmp3568 := PrimTail(tmp3567)

tmp3569 := PrimHead(tmp3568)

tmp3570 := PrimTail(tmp3569)

tmp3571 := PrimHead(tmp3570)

tmp3572 := PrimTail(tmp3571)

tmp3573 := PrimIsPair(tmp3572)

var ifres3526 Obj

if True == tmp3573 {
tmp3559 := PrimTail(V881)

tmp3560 := PrimTail(tmp3559)

tmp3561 := PrimHead(tmp3560)

tmp3562 := PrimTail(tmp3561)

tmp3563 := PrimHead(tmp3562)

tmp3564 := PrimTail(tmp3563)

tmp3565 := PrimTail(tmp3564)

tmp3566 := PrimEqual(Nil, tmp3565)

var ifres3528 Obj

if True == tmp3566 {
tmp3553 := PrimTail(V881)

tmp3554 := PrimTail(tmp3553)

tmp3555 := PrimHead(tmp3554)

tmp3556 := PrimTail(tmp3555)

tmp3557 := PrimTail(tmp3556)

tmp3558 := PrimEqual(Nil, tmp3557)

var ifres3530 Obj

if True == tmp3558 {
tmp3549 := PrimTail(V881)

tmp3550 := PrimTail(tmp3549)

tmp3551 := PrimTail(tmp3550)

tmp3552 := PrimIsPair(tmp3551)

var ifres3532 Obj

if True == tmp3552 {
tmp3544 := PrimTail(V881)

tmp3545 := PrimTail(tmp3544)

tmp3546 := PrimTail(tmp3545)

tmp3547 := PrimTail(tmp3546)

tmp3548 := PrimEqual(Nil, tmp3547)

var ifres3534 Obj

if True == tmp3548 {
tmp3536 := PrimTail(V881)

tmp3537 := PrimTail(tmp3536)

tmp3538 := PrimHead(tmp3537)

tmp3539 := PrimTail(tmp3538)

tmp3540 := PrimHead(tmp3539)

tmp3541 := PrimTail(tmp3540)

tmp3542 := PrimHead(tmp3541)

tmp3543 := PrimIsSymbol(tmp3542)

var ifres3535 Obj

if True == tmp3543 {
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

var ifres3527 Obj

if True == ifres3528 {
ifres3527 = True


} else {
ifres3527 = False


}

ifres3526 = ifres3527


} else {
ifres3526 = False


}

var ifres3525 Obj

if True == ifres3526 {
ifres3525 = True


} else {
ifres3525 = False


}

ifres3524 = ifres3525


} else {
ifres3524 = False


}

var ifres3523 Obj

if True == ifres3524 {
ifres3523 = True


} else {
ifres3523 = False


}

ifres3522 = ifres3523


} else {
ifres3522 = False


}

var ifres3521 Obj

if True == ifres3522 {
ifres3521 = True


} else {
ifres3521 = False


}

ifres3520 = ifres3521


} else {
ifres3520 = False


}

var ifres3519 Obj

if True == ifres3520 {
ifres3519 = True


} else {
ifres3519 = False


}

ifres3518 = ifres3519


} else {
ifres3518 = False


}

var ifres3517 Obj

if True == ifres3518 {
ifres3517 = True


} else {
ifres3517 = False


}

ifres3516 = ifres3517


} else {
ifres3516 = False


}

var ifres3515 Obj

if True == ifres3516 {
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

var ifres3511 Obj

if True == ifres3512 {
ifres3511 = True


} else {
ifres3511 = False


}

ifres3510 = ifres3511


} else {
ifres3510 = False


}

var ifres3509 Obj

if True == ifres3510 {
ifres3509 = True


} else {
ifres3509 = False


}

ifres3508 = ifres3509


} else {
ifres3508 = False


}

if True == ifres3508 {
tmp3494 := PrimTail(V881)

tmp3495 := PrimTail(tmp3494)

tmp3496 := PrimHead(tmp3495)

tmp3497 := PrimTail(tmp3496)

tmp3498 := PrimHead(tmp3497)

tmp3499 := PrimTail(tmp3498)

tmp3500 := PrimHead(tmp3499)

tmp3501 := PrimTail(V881)

tmp3502 := PrimHead(tmp3501)

tmp3503 := PrimTail(V881)

tmp3504 := PrimTail(tmp3503)

tmp3505 := PrimTail(tmp3504)

tmp3506 := PrimHead(tmp3505)

__e.TailApply(PrimFunc(symsubst), tmp3500, tmp3502, tmp3506)
return


} else {
__e.Return(V881)
return
}


}, 1)

tmp3609 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp3493)


_ = tmp3609

tmp3610 := MakeNative(func(__e *ControlFlow) {
V884 := __e.Get(1)
_ = V884
V885 := __e.Get(2)
_ = V885
V886 := __e.Get(3)
_ = V886
tmp3709 := PrimIsPair(V885)

var ifres3655 Obj

if True == tmp3709 {
tmp3707 := PrimHead(V885)

tmp3708 := PrimIsPair(tmp3707)

var ifres3657 Obj

if True == tmp3708 {
tmp3704 := PrimHead(V885)

tmp3705 := PrimHead(tmp3704)

tmp3706 := PrimIsPair(tmp3705)

var ifres3659 Obj

if True == tmp3706 {
tmp3700 := PrimHead(V885)

tmp3701 := PrimHead(tmp3700)

tmp3702 := PrimHead(tmp3701)

tmp3703 := PrimEqual(symand, tmp3702)

var ifres3661 Obj

if True == tmp3703 {
tmp3696 := PrimHead(V885)

tmp3697 := PrimHead(tmp3696)

tmp3698 := PrimTail(tmp3697)

tmp3699 := PrimIsPair(tmp3698)

var ifres3663 Obj

if True == tmp3699 {
tmp3691 := PrimHead(V885)

tmp3692 := PrimHead(tmp3691)

tmp3693 := PrimTail(tmp3692)

tmp3694 := PrimTail(tmp3693)

tmp3695 := PrimIsPair(tmp3694)

var ifres3665 Obj

if True == tmp3695 {
tmp3685 := PrimHead(V885)

tmp3686 := PrimHead(tmp3685)

tmp3687 := PrimTail(tmp3686)

tmp3688 := PrimTail(tmp3687)

tmp3689 := PrimTail(tmp3688)

tmp3690 := PrimEqual(Nil, tmp3689)

var ifres3667 Obj

if True == tmp3690 {
tmp3682 := PrimHead(V885)

tmp3683 := PrimTail(tmp3682)

tmp3684 := PrimIsPair(tmp3683)

var ifres3669 Obj

if True == tmp3684 {
tmp3678 := PrimHead(V885)

tmp3679 := PrimTail(tmp3678)

tmp3680 := PrimTail(tmp3679)

tmp3681 := PrimEqual(Nil, tmp3680)

var ifres3671 Obj

if True == tmp3681 {
tmp3673 := PrimHead(V885)

tmp3674 := PrimHead(tmp3673)

tmp3675 := PrimTail(tmp3674)

tmp3676 := PrimHead(tmp3675)

tmp3677 := PrimEqual(V884, tmp3676)

var ifres3672 Obj

if True == tmp3677 {
ifres3672 = True


} else {
ifres3672 = False


}

ifres3671 = ifres3672


} else {
ifres3671 = False


}

var ifres3670 Obj

if True == ifres3671 {
ifres3670 = True


} else {
ifres3670 = False


}

ifres3669 = ifres3670


} else {
ifres3669 = False


}

var ifres3668 Obj

if True == ifres3669 {
ifres3668 = True


} else {
ifres3668 = False


}

ifres3667 = ifres3668


} else {
ifres3667 = False


}

var ifres3666 Obj

if True == ifres3667 {
ifres3666 = True


} else {
ifres3666 = False


}

ifres3665 = ifres3666


} else {
ifres3665 = False


}

var ifres3664 Obj

if True == ifres3665 {
ifres3664 = True


} else {
ifres3664 = False


}

ifres3663 = ifres3664


} else {
ifres3663 = False


}

var ifres3662 Obj

if True == ifres3663 {
ifres3662 = True


} else {
ifres3662 = False


}

ifres3661 = ifres3662


} else {
ifres3661 = False


}

var ifres3660 Obj

if True == ifres3661 {
ifres3660 = True


} else {
ifres3660 = False


}

ifres3659 = ifres3660


} else {
ifres3659 = False


}

var ifres3658 Obj

if True == ifres3659 {
ifres3658 = True


} else {
ifres3658 = False


}

ifres3657 = ifres3658


} else {
ifres3657 = False


}

var ifres3656 Obj

if True == ifres3657 {
ifres3656 = True


} else {
ifres3656 = False


}

ifres3655 = ifres3656


} else {
ifres3655 = False


}

if True == ifres3655 {
tmp3611 := PrimHead(V885)

tmp3612 := PrimHead(tmp3611)

tmp3613 := PrimTail(tmp3612)

tmp3614 := PrimHead(tmp3613)

tmp3615 := PrimTail(V885)

tmp3616 := PrimHead(V885)

tmp3617 := PrimHead(tmp3616)

tmp3618 := PrimTail(tmp3617)

tmp3619 := PrimTail(tmp3618)

tmp3620 := PrimHead(tmp3619)

tmp3621 := PrimHead(V885)

tmp3622 := PrimTail(tmp3621)

tmp3623 := PrimCons(tmp3620, tmp3622)

tmp3624 := PrimCons(tmp3623, V886)

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp3614, tmp3615, tmp3624)
return


} else {
tmp3653 := PrimIsPair(V885)

var ifres3633 Obj

if True == tmp3653 {
tmp3651 := PrimHead(V885)

tmp3652 := PrimIsPair(tmp3651)

var ifres3635 Obj

if True == tmp3652 {
tmp3648 := PrimHead(V885)

tmp3649 := PrimTail(tmp3648)

tmp3650 := PrimIsPair(tmp3649)

var ifres3637 Obj

if True == tmp3650 {
tmp3644 := PrimHead(V885)

tmp3645 := PrimTail(tmp3644)

tmp3646 := PrimTail(tmp3645)

tmp3647 := PrimEqual(Nil, tmp3646)

var ifres3639 Obj

if True == tmp3647 {
tmp3641 := PrimHead(V885)

tmp3642 := PrimHead(tmp3641)

tmp3643 := PrimEqual(V884, tmp3642)

var ifres3640 Obj

if True == tmp3643 {
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

if True == ifres3633 {
tmp3625 := PrimHead(V885)

tmp3626 := PrimHead(tmp3625)

tmp3627 := PrimTail(V885)

tmp3628 := PrimHead(V885)

tmp3629 := PrimTail(tmp3628)

tmp3630 := PrimCons(True, tmp3629)

tmp3631 := PrimCons(tmp3630, V886)

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp3626, tmp3627, tmp3631)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V886, V885)
return
}


}


}, 3)

tmp3710 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp3610)


_ = tmp3710

tmp3711 := MakeNative(func(__e *ControlFlow) {
V889 := __e.Get(1)
_ = V889
V890 := __e.Get(2)
_ = V890
tmp3735 := PrimIsPair(V889)

var ifres3726 Obj

if True == tmp3735 {
tmp3733 := PrimTail(V889)

tmp3734 := PrimIsPair(tmp3733)

var ifres3728 Obj

if True == tmp3734 {
tmp3730 := PrimTail(V889)

tmp3731 := PrimTail(tmp3730)

tmp3732 := PrimEqual(Nil, tmp3731)

var ifres3729 Obj

if True == tmp3732 {
ifres3729 = True


} else {
ifres3729 = False


}

ifres3728 = ifres3729


} else {
ifres3728 = False


}

var ifres3727 Obj

if True == ifres3728 {
ifres3727 = True


} else {
ifres3727 = False


}

ifres3726 = ifres3727


} else {
ifres3726 = False


}

if True == ifres3726 {
tmp3712 := MakeNative(func(__e *ControlFlow) {
W891 := __e.Get(1)
_ = W891
tmp3722 := PrimEqual(symshen_4skip, W891)

if True == tmp3722 {
__e.Return(V890)
return
} else {
tmp3713 := Call(__e, PrimFunc(symshen_4op1), W891)


tmp3714 := PrimTail(V889)

tmp3715 := PrimCons(tmp3713, tmp3714)

tmp3716 := Call(__e, PrimFunc(symshen_4op2), W891)


tmp3717 := PrimTail(V889)

tmp3718 := PrimCons(tmp3716, tmp3717)

tmp3719 := PrimCons(tmp3718, Nil)

tmp3720 := PrimCons(tmp3715, tmp3719)

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp3720, V890)
return


}


}, 1)

tmp3723 := PrimHead(V889)

tmp3724 := Call(__e, PrimFunc(symshen_4op), tmp3723)


__e.TailApply(tmp3712, tmp3724)
return


} else {
__e.Return(V890)
return
}


}, 2)

tmp3736 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp3711)


_ = tmp3736

tmp3737 := MakeNative(func(__e *ControlFlow) {
V894 := __e.Get(1)
_ = V894
tmp3745 := PrimEqual(symcons_2, V894)

if True == tmp3745 {
__e.Return(symcons)
return
} else {
tmp3743 := PrimEqual(symshen_4_7string_2, V894)

if True == tmp3743 {
__e.Return(sym_8s)
return
} else {
tmp3741 := PrimEqual(symshen_4_7vector_2, V894)

if True == tmp3741 {
__e.Return(sym_8v)
return
} else {
tmp3739 := PrimEqual(symtuple_2, V894)

if True == tmp3739 {
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

tmp3746 := Call(__e, ns2_1set, symshen_4op, tmp3737)


_ = tmp3746

tmp3747 := MakeNative(func(__e *ControlFlow) {
V895 := __e.Get(1)
_ = V895
V896 := __e.Get(2)
_ = V896
tmp3766 := PrimEqual(Nil, V895)

if True == tmp3766 {
__e.Return(V896)
return
} else {
tmp3764 := PrimIsPair(V895)

if True == tmp3764 {
tmp3760 := PrimHead(V895)

tmp3761 := Call(__e, PrimFunc(symoccurrences), tmp3760, V896)


tmp3762 := PrimGreatThan(tmp3761, MakeNumber(1))

if True == tmp3762 {
tmp3748 := MakeNative(func(__e *ControlFlow) {
W897 := __e.Get(1)
_ = W897
tmp3749 := PrimHead(V895)

tmp3750 := PrimTail(V895)

tmp3751 := PrimHead(V895)

tmp3752 := Call(__e, PrimFunc(symsubst), W897, tmp3751, V896)


tmp3753 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp3750, tmp3752)


tmp3754 := PrimCons(tmp3753, Nil)

tmp3755 := PrimCons(tmp3749, tmp3754)

tmp3756 := PrimCons(W897, tmp3755)

__e.Return(PrimCons(symlet, tmp3756))
return


}, 1)

tmp3757 := Call(__e, PrimFunc(symgensym), symSelect)


__e.TailApply(tmp3748, tmp3757)
return


} else {
tmp3758 := PrimTail(V895)

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp3758, V896)
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1selectors_1h)
return
}


}


}, 2)

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp3747)
return




}, 0)

