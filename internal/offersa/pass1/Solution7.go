package pass1

import(
	"sort"
)
/**
 * 剑指 Offer II 007. 数组中和为0的三个数
 * - 数组元素是无序的
 * - 找出所有和为0的三元组
 * - 两个三元组中对应位置元素不能相同
 */

type Solution7 struct{}

    /**
     * 先升序排序
     * - 再使用"剑指 Offer II 006"的方法 双指针反向移动
     * - 不使用set防止三元组重复, 手动控制
     * - 优化 遍历过程中相连的相等元素只使用最后1个，避免两个三元组对应位置元素相等
     * @param nums
     * @return
     */
func (solution7 Solution7) threeSum(nums []int) [][]int{
	sort.Ints(nums)

	var rs[][]int=[][]int{}
	for i:=0;i<len(nums);i+=1{
		// 防止三元组重复 已经处理过的i不再处理
		if i>0 && nums[i]==nums[i-1]{
			continue
		}

		var target int=0-nums[i]
		var j=i+1
		var k=len(nums)-1

		for j<k{
			if nums[j]+nums[k]==target{
				var r []int=[]int{}
				r=append(r,nums[i])
				r=append(r,nums[j])
				r=append(r,nums[k])
				rs=append(rs,r)
				// 已经处理过1次的i，k，不再处理
				for j<k {
                    var numj int= nums[j]
					var numj1 int=nums[j+1]
					if numj!=numj1{
						j++
						break
					}
					j++
				}

				for j<k{
					var numk0 int=nums[k-1]
                    var numk int=nums[k]
					if numk != numk0{
						k--
						break
					}
					k--
				}
			}

			if nums[j]+nums[k]>target{
			    k--
            }
			
			if nums[j]+nums[k]<target{
                j++
            }
		}
	}

	return rs
}
