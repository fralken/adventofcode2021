package day16

import (
	"aoc2021/utils"
	"strconv"
)

func FirstStar() {
	utils.Star(16, 1, "sum of version numbers", firstStar)
}

func SecondStar() {
	utils.Star(16, 2, "result of expression", secondStar)
}

func firstStar(content string) int {
	packets := parsePackets(content)
	return sumVersions(packets)
}

func secondStar(content string) int {
	packets := parsePackets(content)
	results := evals(packets)
	return results[0]
}

type packet struct {
	version int
	typeId int
	value int
	packets []packet
}

func extractPackets(content string, pos int, end int) ([]packet, int) {
	var packets []packet
	for pos < end {
		if pos + 7 < end { // are there enough bits for a packet (version + type + mode)?
			var thisPacket packet
			thisPacket, pos = extractPacket(content, pos)
			packets = append(packets, thisPacket)
		} else {
			pos = end
		}
	}
	return packets, pos
}

func parsePackets(content string) (packets []packet) {
	packets, _ = extractPackets(content, 0, len(content) * 4)
	return
}

func extractPacket(content string, pos int) (packet, int) {
	var thisPacket packet
	var version int
	var typeId int
	version, pos = parseBits(content, pos, 3)
	typeId, pos = parseBits(content, pos, 3)
	if typeId == 4 {
		var value int
		value, pos = parseBitsLiteral(content, pos)
		thisPacket = packet{ version, typeId, value, nil }
	} else {
		var mode int
		mode, pos = parseBits(content, pos, 1)
		if mode == 0 {
			var size int
			var subPackets []packet
			size, pos = parseBits(content, pos, 15)
			subPackets, pos = extractPackets(content, pos, pos + size)
			thisPacket = packet{ version, typeId, 0, subPackets }
		} else { // mode == 1
			var count int
			var subPackets []packet
			count, pos = parseBits(content, pos, 11)
			for i := 0; i < count; i++ {
				var packet packet
				packet, pos = extractPacket(content, pos)
				subPackets = append(subPackets, packet)
			}
			thisPacket = packet{ version, typeId, 0, subPackets }
		}
	}
	return thisPacket, pos
}

func parseBits(content string, pos int, size int) (num int, newPos int) {
	index := pos / 4
	bit := 3 - pos % 4
	digit, _ := strconv.ParseInt(string(content[index]), 16, 0)
	num = 0
	for newPos = pos; size > 0; size-- {
		if bit < 0 {
			bit = 3
			index++
			digit, _ = strconv.ParseInt(string(content[index]), 16, 8)
		}
		num = num << 1 + int((digit) >> bit & 1)
		bit--
		newPos++
	}
	return
}

func parseBitsLiteral(content string, pos int) (int, int) {
	var flag int
	var part int
	num := 0
	for flag, pos = parseBits(content, pos, 1); flag == 1; flag, pos = parseBits(content, pos, 1) {
		part, pos = parseBits(content, pos, 4)
		num = num * 16 + part
	}
	part, pos = parseBits(content, pos, 4)
	num = num * 16 + part
	return num, pos
}

func sumVersions(packets []packet) int {
	sum := 0
	for _, thisPacket := range packets {
		sum += thisPacket.version
		if thisPacket.packets != nil {
			sum += sumVersions(thisPacket.packets)
		}
	}
	return sum
}

func evals(packets []packet) (result []int) {
	for _, p := range packets {
		result = append(result, eval(p))
	}
	return
}

func eval(p packet) int {
	switch p.typeId {
		case 0: { // sum
			sum := 0
			for _, res := range evals(p.packets) {
				sum += res
			}
			return sum
		}
		case 1: { // product
			product := 1
			for _, res := range evals(p.packets) {
				product *= res
			}
			return product
		}
		case 2: { // min
			results := evals(p.packets)
			min := results[0]
			for _, res := range results {
				if res < min { min = res }
			}
			return min
		}
		case 3: { // max
			results := evals(p.packets)
			max := results[0]
			for _, res := range results {
				if res > max { max = res }
			}
			return max
		}
		case 4: return p.value
		case 5: { // greater
			results := evals(p.packets)
			if results[0] > results[1] { return 1 } else { return 0 }
		}
		case 6: { // less
			results := evals(p.packets)
			if results[0] < results[1] { return 1 } else { return 0 }
		}
		case 7: { // equal
			results := evals(p.packets)
			if results[0] == results[1] { return 1 } else { return 0 }
		}
		default:
			return 0
	}
}