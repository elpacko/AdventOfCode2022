package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 01 baby!")
	file, err := os.Open("day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//! Inrease the buffer capacity if necessary
	const maxCapacity = 10_000 * 1024 // 20GB == 20_000*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	var max_calorie_elf_qty int = 0
	var current_calorie_elf_qty int = 0
	var max_calorie_elfs_qty [3]int
	for scanner.Scan() {
		var current_line = scanner.Text()
		if scanner.Text() == "" {
			var elf_id_shift = -1
			for elfid, elf_cal := range max_calorie_elfs_qty {
				if current_calorie_elf_qty >= elf_cal {
					elf_id_shift = elfid
				}
			}
			if elf_id_shift != -1 {
				print(fmt.Sprintf("------%d-----\n", elf_id_shift))
				for i := 0; i < elf_id_shift; i++ {
					max_calorie_elfs_qty[i] = max_calorie_elfs_qty[i+1]
				}
				max_calorie_elfs_qty[elf_id_shift] = current_calorie_elf_qty

				for i := 0; i < 3; i++ {
					print(fmt.Sprintf("%d\t%d\n", max_calorie_elfs_qty[i], i))
				}
			}
			// if current_calorie_elf_qty > max_calorie_elf_qty {
			// 	max_calorie_elf_qty = current_calorie_elf_qty
			// }

			current_calorie_elf_qty = 0
		} else {
			calorie_int, _ := strconv.Atoi(current_line)
			current_calorie_elf_qty = current_calorie_elf_qty + calorie_int
		}

	}
	for elf_i, elf_qty := range max_calorie_elfs_qty {
		print(fmt.Sprintf("%d\t%d\n", elf_i, elf_qty))
		max_calorie_elf_qty += elf_qty
	}
	print("----- total:")
	print(max_calorie_elf_qty)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
