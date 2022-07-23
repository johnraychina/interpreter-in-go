# 用Go语言实现一个解释器interpreter

2022年我学go语言的时候，鬼使神差发现一个电子书《writing an interpreter in go》，评分还挺高的。

想着之前学java的时候没能研究一下JIT，有点遗憾。 
现在能学go又能学编译、解释器的原理，一举两得，看到这个电子书真是缘分。

**经历**

因为没有设定计划，学到哪里算哪里，不会有时间焦虑。

从5月初到7月初，历时两个月，左右边看边写，平常摸鱼1小时写，周末花半天，
主逻辑是手写的，大部分test case是copy的。
因为有对照，而且有很多test，每次做到小步调试，有问题也很快解决了。

**感受**
所以，总的来说学习效率很高，有一种酣畅淋漓的感觉。

作者绝对是个大神，深入浅出，只讲必要的东西，专注语言层面，垃圾回就用go自带的，没有炫技，
整个过程感觉是他用最简明扼要的方式带着我一步步攀登。

首先搭好一个大的架子：
lexer分词，parser解析出来AST(Top down precedence operator算法)，eval运算，object对象体系，Environment保存变量。

然后逐步往里面填充语言特性：简单赋值语句，if条件表达式，函数，返回值，数据结构。

每一步都很扎实，前后衔接非常顺畅， 没有任何跳跃感。

# 下一步？
- Writing a compile in go 解释器后面就是编译器，很自然。
- LLVM 一个在Apple创作了swift语言的大神，觉得编译器*计算机平台 组合这么多，不通用，太浪费人力，做了一个通用的方案，现在似乎是行业标准了。
- 机器学习编译 https://mlc.ai/summer22-zh/ 陈天奇大神，看看机器学习平台怎么做编译的。
