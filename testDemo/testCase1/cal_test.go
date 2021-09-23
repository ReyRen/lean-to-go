package testCase1

import "testing" // 引入testing框架就有个隐藏的main函数

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper execute err, expect=%v but=%v\n", 55, res)
	}

	t.Logf("addUpper execute successful")
}

/*
	testing提供对Go包自动化测试的支持。通过`go test`命令，能够自动执行func TestXxx(test *testing.T)形式的函数
	Xxx可以是任意字符，但是首字母必须大写

	要编写一个新的测试套件，需要创建一个名称以_test.go结尾的文件，该文件包含TestXxx()函数, 将该文件放在与被测试的包相同的包中，
	该文件将被排除在正常的程序包之外，但在运行go test命令时才将包包含。
*/
