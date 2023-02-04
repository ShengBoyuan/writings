package bitarray

// ToInt 接收原子权限 Seq 的切片, 根据位次转化为 BitArray.
// 最终计算 BitArray 对应的 int 列表并返回.
func ToInt(permSeqs []int) []uint64 {
	var bitArray1 uint64
	var bitArray2 uint64
	for _, seq := range permSeqs {
		if seq != 0 {
			if seq <= 64 {
				bitArray1 |= 1 << (seq - 1)
			} else {
				bitArray2 |= 1 << (seq - 64 - 1)
			}
			// TODO 第三次扩展.
		}
	}
	return []uint64{bitArray1, bitArray2}
}

// ToSlice 接收 BitArray 对应的 int 值, 根据位次解析原子权限 Seq.
// 最终以原子权限 Seq 列表的格式返回.
// 数组中 bitArray 权重依 index 升高, 0 位首个 bitArray 代表 1~64 原子权限, 1 位代表 65~128 原子权限.
// 最终返回结果会去除重复 seq, 并且会升序排列.
func ToSlice(bitArrays []uint64) []int {
	var permSeqs []int
	for index, bitArray := range bitArrays {
		var flag uint64 = 1
		for i := 1; i <= 64; i++ {
			if bitArray&flag != 0 {
				permSeqs = append(permSeqs, i+index*64)
			}
			flag = flag << 1
		}
	}
	return permSeqs
}

// Union 接收多个 BitArray 组成的切片,将其去重取并集后返回一个 BitArray.
func Union(bitArrays []uint64) uint64 {
	if len(bitArrays) == 0 {
		return 0
	}
	var permBitArray uint64
	for _, bitArray := range bitArrays {
		permBitArray |= bitArray
	}

	return permBitArray
}

// Intersect 接收多个 BitArray 组成的切片, 将其去重取交集后返回一个 BitArray.
func Intersect(bitArrays []uint64) uint64 {
	if len(bitArrays) == 0 {
		return 0
	}
	var permBitArray uint64 = 0xffffffffffffffff
	for _, bitArray := range bitArrays {
		permBitArray &= bitArray
	}

	return permBitArray
}

// Include 接受两个 BitArray, 判断 a 是否包含 b 的所有原子权限.
// 若全部包含则返回 true, 否则返回 false.
func Include(a uint64, b uint64) bool {
	return b == a&b
}

// Delete 将一个原子权限 BitArray 取出部分权限.
// 接收 BitArray 列表和一个原子权限 Seq 列表.
// 从前者中取出后者.
func Delete(bitArrays []uint64, seqs []int) []uint64 {
	if len(bitArrays) == 0 {
		return bitArrays
	}
	seqBitArrays := ToInt(seqs)
	for index, bitArray := range bitArrays {
		bitArrays[index] = bitArray & ^seqBitArrays[index]
	}
	return bitArrays
}
