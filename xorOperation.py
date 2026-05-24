class Solution:
    def xorOperation(self,n:int,start:int) -> int:
        nums = [0] * n
        for i in range(n):
            nums[i] = start + 2 * i
        result = 0
        for i in range(n):
            result ^= nums[i]
        return result
        

if __name__ == "__main__":
    solution = Solution()
    print(solution.xorOperation(5, 0))
    print(solution.xorOperation(4, 3))
    print(solution.xorOperation(1, 7))
    print(solution.xorOperation(10, 5))
