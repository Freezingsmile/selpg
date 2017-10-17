# selpg

    selpg从标准输入或从作为命令行参数给出的文件名读取文本输入。它允许用户指定来自该输入并随后将被输出的页面范围。
例如，如果输入含有 100 页，则用户可指定只打印第 35 至 65 页。这种特性有实际价值，因为在打印机上打印选定的页面避免了浪费纸张。
另一个示例是，原始文件很大而且以前已打印过，但某些页面由于打印机卡住或其它原因而没有被正确打印。在这样的情况下，则可用该工具
来只打印需要打印的页面。

    selpg 是从文本输入选择页范围的实用程序。该输入可以来自作为最后一个命令行参数指定的文件，在没有给出文件名参数时也可以来自标准输入。
selpg 首先处理所有的命令行参数。在扫描了所有的选项参数（也就是那些以连字符为前缀的参数）后，如果 selpg 发现还有一个参数，则它会接受该参数为输入文件的名称并尝试打开它以进行读取。如果没有其它参数，则 selpg 假定输入来自标准输入。

    参数处理
    
“-sNumber”和“-eNumber”强制选项：
selpg 要求用户用两个命令行参数“-sNumber”（例如，“-s10”表示从第 10 页开始）和“-eNumber”（例如，“-e20”表示在第 20 页结束）指定要抽取的页面范围的起始页和结束页。selpg 对所给的页号进行合理性检查；换句话说，它会检查两个数字是否为有效的正整数以及结束页是否不小于起始页。这两个选项，“-sNumber”和“-eNumber”是强制性的，而且必须是命令行上在命令名 selpg 之后的头两个参数：

$ selpg -s10 -e20 ...

    “-lNumber”和“-f”可选选项：

selpg 可以处理两种输入文本：
类型 1：该类文本的页行数固定。这是缺省类型，因此不必给出选项进行说明。也就是说，如果既没有给出“-lNumber”也没有给出“-f”选项，则 selpg 会理解为页有固定的长度（每页 72 行）。
选择 72 作为缺省值是因为在行打印机上这是很常见的页长度。这样做的意图是将最常见的命令用法作为缺省值，这样用户就不必输入多余的选项。该缺省值可以用“-lNumber”选项覆盖，如下所示：

$ selpg -s10 -e20 -l66 ...

这表明页有固定长度，每页为 66 行。

类型 2：该类型文本的页由 ASCII 换页字符（十进制数值为 12，在 C 中用“\f”表示）定界。该格式与“每页行数固定”格式相比的好处在于，当每页的行数有很大不同而且文件有很多页时，该格式可以节省磁盘空间。在含有文本的行后面，类型 2 的页只需要一个字符 ― 换页 ― 就可以表示该页的结束。打印机会识别换页符并自动根据在新的页开始新行所需的行数移动打印头。
将这一点与类型 1 比较：在类型 1 中，文件必须包含 PAGELEN - CURRENTPAGELEN 个新的行以将文本移至下一页，在这里 PAGELEN 是固定的页大小而 CURRENTPAGELEN 是当前页上实际文本行的数目。在此情况下，为了使打印头移至下一页的页首，打印机实际上必须打印许多新行。这在磁盘空间利用和打印机速度方面效率都很低（尽管实际的区别可能不太大）。
类型 2 格式由“-f”选项表示，如下所示：

$ selpg -s10 -e20 -f ...

该命令告诉 selpg 在输入中寻找换页符，并将其作为页定界符处理。
注：“-lNumber”和“-f”选项是互斥的。

    “-dDestination”可选选项：
selpg 还允许用户使用“-dDestination”选项将选定的页直接发送至打印机。这里，“Destination”应该是 lp 命令“-d”选项（请参阅“man lp”）可接受的打印目的地名称。该目的地应该存在 ― selpg 不检查这一点。在运行了带“-d”选项的 selpg 命令后，若要验证该选项是否已生效，请运行命令“lpstat -t”。该命令应该显示添加到“Destination”打印队列的一项打印作业。如果当前有打印机连接至该目的地并且是启用的，则打印机应打印该输出。这一特性是用 popen() 系统调用实现的，该系统调用允许一个进程打开到另一个进程的管道，将管道用于输出或输入。在下面的示例中，我们打开到命令

$ lp -dDestination

的管道以便输出，并写至该管道而不是标准输出：

selpg -s10 -e20 -dlp1

该命令将选定的页作为打印作业发送至 lp1 打印目的地。应该可以看到类似“request id is lp1-6”的消息。该消息来自 lp 命令；它显示打印作业标识。如果在运行 selpg 命令之后立即运行命令 
lpstat -t | grep lp1 ，
应该看见 lp1 队列中的作业。如果在运行 lpstat 命令前耽搁了一些时间，那么可能看不到该作业，因为它一旦被打印就从队列中消失了。
 


测试 
 
使用 txt文本进行测试:

The day that you see me old and I am already not, have patience and try to understand me .
If I get dirty when eating… if I cannot dress… have patience.
Remember the hours I spent teaching it to you.
If, when I speak to you, I repeat the same things  thousand and one times.
Do not interrupt me. Listen to me.
When you were small, I had to read to you thousand and one times the same story 
until you get to sleep… When I do not want to have a shower, neither shame me nor 
scold me. Remember when I had to chase you with thousand excuses I invented, 
in order that you wanted 
to bath…When you see my on new technologies… give me the necessary time and not look at me 
with your mocking smile.
I taught you how to do so many things… to eat well, to dress well… to confront life.
When at some moment I lose the memory or the thread of our conversation… let me have the necessary time 
to remember… and if I cannot do it；do not become nervous… as the most important thing is not 
my conversation but surely to be with you and to have you listening to me.
If ever I do not want to eat, do not force me. I know well when I need to and when not.
When my tired legs do not allow me walk… give me your hand… the same way 
I did when you gave your first steps. 
And when someday I say to you that I do not want to live any more.
that I want to die… do not get angry… some day you will understand.
Try to understand that my age is not lived but survived. 
Some day you will discover that, despite my mistakes, I always wanted 
the best thing for you and that I tried to prepare the way for you. 
You must not feel sad, angry or impotent for seeing me near you. 
You must be next to me, 
try to understand me and to help me as I did it when you started living. 
Help me to walk.
Help me to end my way with love and patience. 
I will pay you by a smile and by the immense love I have had always for you. 
I love you son.

    Test cat test | selpg -s 2 -e 3 -l 5
result : 

with your mocking smile.
I taught you how to do so many things… to eat well, to dress well… to confront life.
When at some moment I lose the memory or the thread of our conversation… let me have the necessary time 
to remember… and if I cannot do it；do not become nervous… as the most important thing is not 
my conversation but surely to be with you and to have you listening to me.
If ever I do not want to eat, do not force me. I know well when I need to and when not.
When my tired legs do not allow me walk… give me your hand… the same way 
I did when you gave your first steps. 
And when someday I say to you that I do not want to live any more.
that I want to die… do not get angry… some day you will understand.

    Test selpg -s 1 -e 1 -l 20 test
result:

the best thing for you and that I tried to prepare the way for you. 
You must not feel sad, angry or impotent for seeing me near you. 
You must be next to me, 
try to understand me and to help me as I did it when you started living. 
Help me to walk.
Help me to end my way with love and patience. 
I will pay you by a smile and by the immense love I have had always for you. 
I love you son.

    Test selpg -s=1 -e=1 -l 6 < test
result:

until you get to sleep… When I do not want to have a shower, neither shame me nor 
scold me. Remember when I had to chase you with thousand excuses I invented, 
in order that you wanted 
to bath…When you see my on new technologies… give me the necessary time and not look at me 
with your mocking smile.
I taught you how to do so many things… to eat well, to dress well… to confront life.
When at some moment I lose the memory or the thread of our conversation… let me have the necessary time 
to remember… and if I cannot do it；do not become nervous… as the most important thing is not 
my conversation but surely to be with you and to have you listening to me.
If ever I do not want to eat, do not force me. I know well when I need to and when not.
When my tired legs do not allow me walk… give me your hand… the same way 
I did when you gave your first steps. 

    Test selpg -s 2 -e 2 -l 5 test >output
result:

with your mocking smile.
I taught you how to do so many things… to eat well, to dress well… to confront life.
When at some moment I lose the memory or the thread of our conversation… let me have the necessary time 
to remember… and if I cannot do it；do not become nervous… as the most important thing is not 
my conversation but surely to be with you and to have you listening to me.

    Test ./selpg -s 1 -e 2 -l 2 cat -n printer test
result:

Remember the hours I spent teaching it to you.
If, when I speak to you, I repeat the same things  thousand and one times.
Do not interrupt me. Listen to me.
When you were small, I had to read to you thousand and one times the same story 
（ 这里本意是测试-d命令，不过因为没有设备，所以只能用cat -n)

    Test ./selpg -s 1 -e 2 -l 2 test | go run scanner.go
使用 scanner.go 文件中的函数（文件可在附件中找到）
result:
Scanner: Remember the hours I spent teaching it to you.
Scanner:If, when I speak to you, I repeat the same things  thousand and one times.
Scanner:Do not interrupt me. Listen to me.
Scanner:When you were small, I had to read to you thousand and one times the same story 

 
以上测试结果均没有出现错误.
