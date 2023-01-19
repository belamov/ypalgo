package main

//https://contest.yandex.ru/contest/23638/problems/K/

// Гоше дали задание написать красивую сортировку слиянием.
//Поэтому Гоше обязательно надо реализовать отдельно функцию merge и функцию merge_sort.
//Функция merge принимает два отсортированных массива, сливает их в один отсортированный
//массив и возвращает его. Если требуемая сигнатура имеет вид merge(array, left, mid, right), то
//первый массив задаётся полуинтервалом [left,mid) массива array, а второй – полуинтервалом
//[mid,right) массива array.
//Функция merge_sort принимает некоторый подмассив, который нужно отсортировать.
//Подмассив задаётся полуинтервалом — его началом и концом.
//Функция должна отсортировать передаваемый в неё подмассив, она ничего не возвращает.
//Функция merge_sort разбивает полуинтервал на две половинки и рекурсивно
//вызывает сортировку отдельно для каждой.
//Затем два отсортированных массива сливаются в один с помощью merge.
//Заметьте, что в функции передаются именно полуинтервалы [begin,end), то есть
//правый конец не включается. Например, если вызвать merge_sort(arr, 0, 4), где
//arr = [4,5,3,0,1,2], то будут отсортированы только первые четыре элемента, изменённый
//массив будет выглядеть как arr=[0,3,4,5,1,2]
//Реализуйте эти две функции.

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf == 1 {
		return
	}

	mid := (lf + rg) / 2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)

	res := merge(arr, lf, mid, rg)
	for i := lf; i < rg; i++ {
		arr[i] = res[i-lf]
	}
}

func merge(arr []int, left int, mid int, right int) []int {
	l, r := left, mid
	result := make([]int, 0, right-left)
	for l < mid && r < right {
		if arr[l] <= arr[r] {
			result = append(result, arr[l])
			l++
		} else {
			result = append(result, arr[r])
			r++
		}
	}

	for l < mid {
		result = append(result, arr[l])
		l++
	}
	for r < right {
		result = append(result, arr[r])
		r++
	}

	return result
}
