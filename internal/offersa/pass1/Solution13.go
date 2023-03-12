package pass1

/**
 * 剑指 Offer II 013. 二维子矩阵的和
 * - 子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)
 */

type NumMatrix1 struct {
	//xcrj 只是声明
	preSum [][]int
}
    /**
     * 一维前缀和，一行的和
     */
// xcrj 工厂模式创建类，静态工厂
// (numMatrix1 *NumMatrix1) 或 (numMatrix1 NumMatrix1)
func Constructor(matrix [][]int) NumMatrix1 {
	//xcrj 二维slice 空间分配
	var numRow int=len(matrix)
	var numCol int=len(matrix[0])
	var preSum [][]int=make([][]int,numRow)
	//xcrj 只写一个是索引
	for i:=range preSum{
		preSum[i]=make([]int,numCol+1)
	}
	// 初始化
	for i:=0;i<numRow;i+=1{
		preSum[i][0]=0
	}
	//构建
	for i:=0;i<numRow;i+=1{
		for j:=0;j<numCol;j+=1{
			preSum[i][j+1]=preSum[i][j]+matrix[i][j]
		}
	}
	
	//xcrj 创建完整对象，返回对象非返回指针
	return NumMatrix1{preSum}
}

func (numMatrix1 *NumMatrix1) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// xcrj i++
	var r int=0
	for i:=row1;i<=row2;i++{
		r+=numMatrix1.preSum[i][col2+1]-numMatrix1.preSum[i][col1]
	}
	return r
}

type NumMatrix2 struct {
	preSum [][]int
}

func Constructor2(matrix [][]int) NumMatrix2 {
	//定义
	var numRow int=len(matrix)
	var numCol int=len(matrix[0])
	var preSum [][]int=make([][]int,numRow+1)
	for i:=0;i<numRow+1;i++{
		preSum[i]=make([]int,numCol+1)
	}
	//初始化
	preSum[0][0]=0
	//构造
	for i:=0;i<numRow;i++{
		for j:=0;j<numCol;j++{
			preSum[i+1][j+1]=preSum[i][j+1]+preSum[i+1][j]-preSum[i][j]+matrix[i][j]
		}
	}

	return NumMatrix2{preSum}
}

func (numMatrix2 *NumMatrix2) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return numMatrix2.preSum[row2+1][col2+1]-numMatrix2.preSum[row2+1][col1]-numMatrix2.preSum[row1][col2+1]+numMatrix2.preSum[row1][col1]
}