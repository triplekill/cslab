# coding:utf8


class Solution(object):
    """矩阵转置
    """

    def transpose(self, A):
        """
        :type A: List[List[int]]
        :rtype: List[List[int]]

        >>> Solution().transpose([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
        [[1, 4, 7], [2, 5, 8], [3, 6, 9]]
        >>> Solution().transpose([[1, 2, 3], [4, 5, 6]])
        [[1, 4], [2, 5], [3, 6]]

        """
        return [[row[col] for row in A] for col in range(len(A[0]))]


def main():
    print(Solution().transpose([[1, 2, 3], [4, 5, 6], [7, 8, 9]]))
    print(Solution().transpose([[1, 2, 3], [4, 5, 6]]))


if __name__ == '__main__':
    main()
