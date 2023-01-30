28. 找出字符串中第一个匹配项的下标
给你两个字符串`haystack`和`needle`，请你在`haystack`字符串中找出`needle`字符串的第一个匹配项的下标（下标从 0 开始）。如果`needle`不是`haystack`的一部分，则返回  -1 。

示例 1：

>输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。

示例 2：
>输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

### 什么是KMP
说到KMP，先说一下KMP这个名字是怎么来的，为什么叫做KMP呢。
因为是由这三位学者发明的：Knuth，Morris和Pratt，所以取了三位学者名字的首字母。所以叫做KMP

### KMP有什么用
KMP主要应用在字符串匹配上。

KMP的主要思想是当出现字符串不匹配时，可以知道一部分之前已经匹配的文本内容，可以利用这些信息避免从头再去做匹配了。

